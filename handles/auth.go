package handlers

import (
	"context"
	"crypto/rand"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
	"gopkg.in/gomail.v2"
	"lms/models"
	"lms/repo"
	"lms/structs"
	"lms/utilS"
	"log"
	"math/big"
	"os"
	"time"
)

var SessAuth = session.New(session.Config{
	//Expiration: 3600,
	CookieSessionOnly: true,
})

func SendEmailToActivations(to string, subject, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "noreply.bitcare@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "noreply.bitcare@gmail.com", "dlmd wcie mrnv ectj")
	return dialer.DialAndSend(mailer)
}

func GetLogin(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLogin")
	sess, _ := SessAuth.Get(c)
	result := sess.Get("authentication")
	if result != nil {
		return c.Redirect("/dashboard")
	} else {
		return c.Render("pages/auth/login", nil)
	}
}

func GetSignup(c *fiber.Ctx) error {
	repo.OutPutDebug(c.IP() + "GetSignup")
	sess, _ := SessAuth.Get(c)
	result := sess.Get("authentication")

	if result != nil {
		return c.Redirect("/dashboard")
	} else {
		return c.Render("pages/auth/signup", nil)
	}
}

func ApiPostLogin(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiPostLogin")
	var form structs.FormLogin

	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNoContent, "Incorrect format", nil)
	}

	if user, err := utilS.UserRepo.FindByEmail(form.Email); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Email not found", nil)
	} else {
		if utilS.CheckPasswordHash(form.Password, user.Vault, user.Password) {
			permission, errP := utilS.RolePerRepo.FindByRoleID(user.RoleID)
			if errP != nil {
				repo.OutPutDebugError(errP.Error())
			}

			sess, _ := SessAuth.Get(c)
			sess.Set("authentication", "success")
			sess.Set("user", models.UserWithoutPass{
				UserID:      user.UserID,
				RoleID:      user.RoleID,
				Name:        user.Name,
				Email:       user.Email,
				TypeUserID:  user.TypeUserID,
				IsActivated: user.IsActivated,
			})
			sess.Set("permission", &permission)

			if err := sess.Save(); err != nil {
				repo.OutPutDebugError("Can not save session user")
			}

			return c.JSON(&utilS.Response{
				Code:    fiber.StatusOK,
				Message: "Login successfully",
				Data:    nil,
			})

		} else {
			return utilS.ResultResponse(c, fiber.StatusUnauthorized, "Password incorrect", nil)
		}

	}

}

func ApiPostForgetPassword(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiPostForgetPassword")
	var form structs.FormLogin

	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNoContent, "Incorrect format", nil)
	}

	if user, err := utilS.UserRepo.FindByEmail(form.Email); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Email not found", nil)
	} else {
		token, _ := utilS.GenerateRandomString(128)
		user.Token = token
		user.TokenExpiry = time.Now().Add(time.Minute * 60)
		if err = utilS.UserRepo.Update(*user); err != nil {
			repo.OutPutDebugError(err.Error())
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot update new token", nil)
		}
		utilS.SendEmailReqResetPassword(token, user)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Forget password successfully. Please check your email for a new password.", nil)
}

func ApiPostActivationPassword(c *fiber.Ctx) error {
	userID := utilS.StringToUint32(c.Params("id"))
	token := c.Params("token")

	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "User not found", nil)
	}

	if user.TokenExpiry.Before(time.Now()) {
		return utilS.ResultResponse(c, fiber.StatusForbidden, "Token expired", nil)
	}

	if token != user.Token {
		return utilS.ResultResponse(c, fiber.StatusForbidden, "Token does not match", nil)
	}

	user.Token = ""
	user.IsActivated = true
	user.TokenExpiry = time.Now()

	newPass, _ := utilS.GenerateRandomString(10)
	user.Password, user.Vault, err = utilS.HashPassword(newPass)
	if err := utilS.UserRepo.Update(*user); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot create new password", nil)
	}
	utilS.SendEmailForgetPassword(newPass, user)

	return c.Render("pages/auth/forgot-password", nil)

	//return utilS.ResultResponse(c, fiber.StatusOK, "Forget password successfully. Please check your email for a new password.", nil)
}

func ApiPostSignup(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiPostSignup")
	var form structs.FormSignup

	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNoContent, "Incorrect format", nil)
	}

	if _, err := utilS.UserRepo.FindByEmailWithoutPass(form.Email); err != nil {
		repo.OutPutDebugError(err.Error())
		var newUser models.User
		newUser.Email = form.Email
		newUser.Name = form.Name
		newUser.Password, newUser.Vault, err = utilS.HashPassword(form.Password)
		newUser.TypeUserID = 4
		newUser.RoleID = 4
		newUser.IsActivated = false
		newUser.Token = ""
		newUser.TokenExpiry = time.Now()
		newUser.CreatedAt = time.Now()
		//newUser.CreatedBy = newUser.UserID
		if err := utilS.UserRepo.Create(newUser); err != nil {
			repo.OutPutDebugError(err.Error())
			return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Create new user failed", nil)
		}

		user, _ := utilS.UserRepo.FindByEmailWithoutPass(form.Email)
		return c.JSON(&utilS.Response{
			Code:    fiber.StatusOK,
			Message: "success",
			Data: fiber.Map{
				"user": user,
			},
		})
	} else {
		return utilS.ResultResponse(c, fiber.StatusFound, "Email already exists", nil)
	}
}

func ApiPostActivation(c *fiber.Ctx) error {
	userID := GetInfoUser(c).UserID

	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "User not found", nil)
	}

	if user.IsActivated == true {
		return utilS.ResultResponse(c, fiber.StatusForbidden, "User is already active", nil)
	}

	token, _ := utilS.GenerateRandomString(128)
	id := utilS.Uint32ToString(userID)

	user.Token = token
	user.TokenExpiry = time.Now().Add(time.Minute * 60)

	domainEmail := os.Getenv("DOMAIN_EMAIL")
	title := "Action Required to Activate Your Account"
	body := `
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			.email-container {
				font-family: Arial, sans-serif;
				color: #333;
				line-height: 1.6;
				padding: 20px;
				max-width: 600px;
				background-color: #f9f9f9;
			}
			.btn-activate {
				display: inline-block;
				padding: 10px 20px;
				margin-top: 20px;
				color: #fff !important;
				background-color: #007bff;
				text-decoration: none;
				border-radius: 5px;
			}
			.btn-activate:hover {
				background-color: #0056b3;
			}
		</style>
	</head>
	<body>
		<div class="email-container">
			<h3>LEARNING MANAGEMENT SYSTEM</h3>
			<h2>Activate Your Account</h2>
			<p>Hello,</p>
			<p>Thank you for registering. Please click the button below to activate your account:</p>
			<a href="` + domainEmail + `/api/v1/resend-activation/` + id + `/token/` + token + `" class="btn-activate">Activate Account</a>
			<p>If you did not register for this account, please ignore this email.</p>
			<p>Best regards,<br>The Support Team</p>
		</div>
	</body>
	</html>
	`
	go func() {
		if err := SendEmailToActivations(user.Email, title, body); err != nil {
		}
	}()

	if err := utilS.UserRepo.Update(*user); err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Can not update", nil)
	}

	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Send token successfully",
		Data:    nil,
	})
}

func CheckToken(c *fiber.Ctx) error {
	userID := utilS.StringToUint32(c.Params("id"))
	token := c.Params("token")

	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "User not found", nil)
	}

	if user.TokenExpiry.Before(time.Now()) {
		return utilS.ResultResponse(c, fiber.StatusForbidden, "Token expired", nil)
	}

	if token != user.Token {
		return utilS.ResultResponse(c, fiber.StatusForbidden, "Token does not match", nil)
	}

	user.Token = ""
	user.IsActivated = true
	user.TokenExpiry = time.Now()

	if err := utilS.UserRepo.Update(*user); err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Can not update", nil)
	}

	sess, _ := SessAuth.Get(c)
	sess.Set("user", models.UserWithoutPass{
		UserID:      user.UserID,
		RoleID:      user.RoleID,
		Name:        user.Name,
		Email:       user.Email,
		TypeUserID:  user.TypeUserID,
		IsActivated: user.IsActivated,
	})
	if err := sess.Save(); err != nil {
		repo.OutPutDebugError("Can not save session user")
	}
	return c.Redirect("/dashboard")
}

func CheckPermission(permissionName string, c *fiber.Ctx) bool {
	for _, per := range GetPermission(c) {
		if per.Permission.Permission == permissionName {
			return true
		}
	}
	return false
}

func CheckAuthentication(c *fiber.Ctx) bool {
	sess, _ := SessAuth.Get(c)
	return sess.Get("authentication") != nil
}

func GetLogout(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLogout")
	return c.Redirect("/")
}

func ApiPostLogout(c *fiber.Ctx) error {
	sess, _ := SessAuth.Get(c)

	if sess.Get("authentication") == nil {
		return utilS.ResultResponse(c, fiber.StatusUnauthorized, "Failed to get session", nil)
	}
	if err := sess.Destroy(); err != nil {
		repo.OutPutDebugError("Failed to destroy session: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to log out", nil)
	}

	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Logged out successfully",
		Data:    nil,
	})
}

//
//func CourseStatusMiddleware(c *fiber.Ctx) error {
//
//	courseID := utilS.StringToUint32(c.Params("id"))
//
//	course, err := utilS.CourseRepo.FindByID(courseID)
//	if err != nil {
//
//		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
//	}
//	if !course.Status {
//		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
//	}
//	return c.Next()
//}
//
//func LessonMiddleware(c *fiber.Ctx) error {
//
//	lessonID := utilS.StringToUint32(c.Params("id"))
//
//	lesson, err := utilS.LessonRepo.FindByID(lessonID)
//	if err != nil {
//		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
//	}
//
//	course, err := utilS.CourseRepo.FindByID(lesson.CourseID)
//	if err != nil {
//		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
//	}
//	if !course.Status {
//		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
//	}
//
//	return c.Next()
//}

// login with google
// Bước 1: Chuyển hướng đến Google
func GoogleLogin(c *fiber.Ctx) error {
	url := GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

// Bước 2: Xử lý phản hồi từ Google
func GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Code not found"})
	}

	// Lấy token từ code
	token, err := GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	// Lấy thông tin user từ Google
	userInfo, err := GetGoogleUserInfo(token)
	if err != nil {
		log.Println("Error fetching user info:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}

	// Lấy thông tin email từ Google
	email, _ := userInfo["email"].(string)
	name, _ := userInfo["name"].(string)
	//avatar, _ := userInfo["picture"].(string)

	// Kiểm tra xem user đã tồn tại chưa
	user, err := utilS.UserRepo.FindByEmail(email)
	var newUser models.User
	if err != nil {
		// Tạo user mới
		newUser = models.User{
			Name:        name,
			Email:       email,
			TypeUserID:  4,
			RoleID:      4,
			IsActivated: false,
			TokenExpiry: time.Now(),
			Token:       "",
			CreatedAt:   time.Now(),
			TypeAccount: "SSO",
		}
		if err = utilS.UserRepo.Create(newUser); err != nil {
			return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create user", nil)
		}
	}
	permission, errP := utilS.RolePerRepo.FindByRoleID(newUser.RoleID)
	if errP != nil {
		repo.OutPutDebugError(errP.Error())
	}

	user, err = utilS.UserRepo.FindByEmail(email)
	if err != nil {
		repo.OutPutDebugError(errP.Error())
	}

	sess, _ := SessAuth.Get(c)
	sess.Set("authentication", "success")
	sess.Set("user", models.UserWithoutPass{
		UserID:      user.UserID,
		RoleID:      user.RoleID,
		Name:        user.Name,
		Email:       user.Email,
		TypeUserID:  user.TypeUserID,
		IsActivated: user.IsActivated,
		TypeAccount: user.TypeAccount,
	})
	sess.Set("permission", &permission)

	if err := sess.Save(); err != nil {
		repo.OutPutDebugError("Can not save session user")
	}

	return c.Redirect("/")
}

func RandomData(len int) string {
	randomString, err := utilS.GenerateRandomString(len)
	if err != nil {
		repo.OutPutDebugError("Error generating random string")
	}
	return randomString
}

func randomChoice(choices []string) string {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(choices))))
	return choices[n.Int64()]
}

func randomFullName() string {
	lastNames := []string{"Nguyễn", "Trần", "Lê", "Phạm", "Hoàng", "Phan", "Vũ", "Đặng", "Bùi", "Đỗ"}
	middleNames := []string{"Văn", "Hữu", "Thị", "Gia", "Hoàng", "Minh", "Khánh", "Nhật", "Quang", "Thanh"}
	firstNames := []string{"An", "Bình", "Cường", "Duy", "Hà", "Linh", "Minh", "Nam", "Phong", "Quân", "Trang", "Vy"}

	lastName := randomChoice(lastNames)     // Họ
	middleName := randomChoice(middleNames) // Họ đệm
	firstName := randomChoice(firstNames)   // Tên

	return lastName + " " + middleName + " " + firstName
}

func FakeDataUser(c *fiber.Ctx) error {
	var users []*models.User
	password, vault, err := utilS.HashPassword(RandomData(8))
	if err != nil {
		repo.OutPutDebugError("Error hashing password")
	}
	for i := 0; i < 100; i++ {
		users = append(users, &models.User{
			TypeUserID:  4,
			RoleID:      4,
			Name:        randomFullName(),
			Email:       RandomData(5) + "@gmail.com",
			Vault:       vault,
			IsActivated: true,
			TypeAccount: "",
			Password:    password,
			CreatedBy:   0,
			UpdatedBy:   0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	if err := utilS.UserRepo.CreateAny(users); err != nil {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"success": false,
			"message": "failed to create user",
		})
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "create user success", nil)
}
