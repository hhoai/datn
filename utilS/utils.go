package utilS

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"lms/models"
	"lms/repo"
	"lms/structs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	_ "time/tzdata"
)

var (
	UserRepo       repo.UserRepository
	RolePerRepo    repo.RolePerRepository
	TypeUserRepo   repo.TypeUserRepository
	RoleRepo       repo.RoleRepository
	PermissionRepo repo.PermissionRepository
	DomainEmail    string
)

func init() {

	UserRepo = repo.NewUserRepository()
	RolePerRepo = repo.NewRolePerRepository()
	TypeUserRepo = repo.NewTypeUserRepository()
	RoleRepo = repo.NewRoleRepository()
	PermissionRepo = repo.NewPermissionRepository()

}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResultResponse(ctx *fiber.Ctx, code int, message string, data interface{}) error {
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return ctx.Status(code).JSON(response)
}

func PageErrorResponse(ctx *fiber.Ctx, code int) error {
	if err := ctx.Status(code).SendFile(fmt.Sprintf("./views/errors/%d.html", code)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return nil
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return ResultResponse(ctx, code, err.Error(), nil)
}

func PageErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return PageErrorResponse(ctx, code)
}

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func XORPassword(password, key string) (string, error) {

	xorBytes := make([]byte, len(password))
	for i := 0; i < len(password); i++ {
		xorBytes[i] = password[i] ^ key[i]
	}
	return string(xorBytes), nil
}

func HashPassword(password string) (string, string, error) {
	key, err := GenerateRandomString(10)
	if err != nil {
		return "", "", err
	}

	xorPassword, errX := XORPassword(password, key)
	if errX != nil {
		return "", "", errX
	}

	hashedBytes, errH := bcrypt.GenerateFromPassword([]byte(xorPassword), bcrypt.DefaultCost)
	if errH != nil {
		return "", "", errH
	}

	return string(hashedBytes), key, nil
}

func CheckPasswordHash(password, key, hashedPassword string) bool {
	xorPassword, err := XORPassword(password, key)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(xorPassword))
	return err == nil
}

func StringToUint32(val string) uint32 {
	result, _ := strconv.Atoi(val)
	return uint32(result)
}

func Uint32ToString(val uint32) string {
	return strconv.Itoa(int(val))
}

func GenerateCourseCode(length int) string {
	courseCode, err := GenerateRandomString(length)
	if err != nil {
		return ""
	}
	return strings.ToUpper(courseCode)
}

func GetFilename(path string) string {
	return filepath.Base(path)
}

func CreateFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func StringToTime(str string) time.Time {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	result, err2 := time.ParseInLocation("2006-01-02T15:04", str, location)
	if err2 != nil {
		repo.OutPutDebugError(err2.Error())
	}

	return result
}

func GetDomainEmail() string {
	domains := strings.Split(os.Getenv("DOMAIN"), ",")
	if domains[0] == "" {
		return "127.0.0.1"
	}
	return domains[0]
}

func sendEmail(to string, subject, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "noreply.bitcare@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "noreply.bitcare@gmail.com", "dlmd wcie mrnv ectj")
	return dialer.DialAndSend(mailer)
}

func BodyEmailNewCourse(courseCode string, userID string) string {
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
			<h2>Join New Course</h2>
			<p>Hello,</p>
			<p>Please click the button below to join the new course</p>
			<p>Course code: ` + courseCode + `</p>
			<a href="` + DomainEmail + `/api/v1/request-course/` + userID + `/` + courseCode + `" class="btn-activate">Join New Course</a>
			<p>If you are not taking this course, please ignore this email.</p>
			<p>Best regards,<br>The Support Team</p>
		</div>
	</body>
	</html>
	`
	return body
}

func SendEmailNewCourse(courseCode string, users []*models.UserWithoutPass) {
	title := "Action required to join the new course"

	for _, user := range users {
		go func(user *models.UserWithoutPass, courseCode string, title string) {
			if err := sendEmail(user.Email, title, BodyEmailNewCourse(courseCode, Uint32ToString(user.UserID))); err != nil {
			}
		}(user, courseCode, title)
	}
}

func SendEmailReqResetPassword(token string, user *models.User) {
	title := "Reset Password"
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
			<h2>Reset Password</h2>
			<p>Hello, ` + user.Name + `</p>
			<p>Please click the button below to get new password</p>
			<a href="` + DomainEmail + `/reset-password/` + Uint32ToString(user.UserID) + `/` + token + `" class="btn-activate">Get New Password</a>
			<p>Best regards,<br>The Support Team</p>
		</div>
	</body>
	</html>
	`
	if err := sendEmail(user.Email, title, body); err != nil {
	}
}

func SendEmailForgetPassword(password string, user *models.User) {
	title := "Reset Password"
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
			<h2>Reset Password</h2>
			<p>Hello, ` + user.Name + `</p>
			<p>Please click the button below to login</p>
			<p>New password: ` + password + `</p>
			<a href="` + DomainEmail + `/login" class="btn-activate">Login</a>
			<p>Best regards,<br>The Support Team</p>
		</div>
	</body>
	</html>
	`
	if err := sendEmail(user.Email, title, body); err != nil {
	}
}

func SendEmailReqUpdateEmail(token string, newEmail string, user *models.User) error {
	title := "Change Email"
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
			<h2>Reset Password</h2>
			<p>Hello, ` + user.Name + `</p>
			<p>Please click the button below to change your email to: ` + newEmail + `!</p>
			<a href="` + DomainEmail + `/update-email/` + Uint32ToString(user.UserID) + `/` + token + `/` + newEmail + `" class="btn-activate">Change Email</a>
			<p>Best regards,<br>The Support Team</p>
		</div>
	</body>
	</html>
	`
	if err := sendEmail(user.Email, title, body); err != nil {
		return err
	}
	return nil
}

func CompareAnswers(correctAnswers, userAnswers []*models.TopicQuestionResponse) ([]structs.AnswerResult, uint32) {
	var results []structs.AnswerResult

	countCorrect := 0
	for i := 0; i < len(correctAnswers); i++ {
		correct := correctAnswers[i]
		user := userAnswers[i]

		isCorrect := true
		var options []models.OptionWithoutIsCorrect

		for j := 0; j < len(correct.Options); j++ {
			correctOption := correct.Options[j]
			userOption := user.Options[j]

			if userOption.IsCorrect != correctOption.IsCorrect {
				isCorrect = false
			}

			options = append(options, models.OptionWithoutIsCorrect{
				OptionID:  correctOption.OptionID,
				Content:   correctOption.Content,
				IsCorrect: userOption.IsCorrect,
			})
		}

		if isCorrect == true {
			countCorrect++
		}

		results = append(results, structs.AnswerResult{
			TopicQuestionID: correct.TopicQuestionID,
			IsCorrect:       isCorrect,
			Options:         options,
			Score:           correct.Score,
			QuestionID:      correct.QuestionID,
			TopicID:         correct.TopicID,
			TypeQuestionID:  correct.TypeQuestionID,
		})
	}

	return results, uint32(countCorrect)
}
