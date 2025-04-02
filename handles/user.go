package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"time"
)

func GetInfoUser(c *fiber.Ctx) models.UserWithoutPass {
	sess, _ := SessAuth.Get(c)
	if result, ok := sess.Get("user").(models.UserWithoutPass); ok {
		return result
	}
	repo.OutPutDebugError("Failed to cast session data to user info")
	return models.UserWithoutPass{}
}
func GetUsers(c *fiber.Ctx) error {
	repo.OutPutDebug("GetUsers")
	return c.Render("pages/users/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}
func ApiGetUsers(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetUsers")
	return utilS.ResultResponse(c, fiber.StatusOK, "Users retrieved successfully", utilS.UserRepo.FindAll())
}

func CheckActivateAccount(c *fiber.Ctx) bool {
	return GetInfoUser(c).IsActivated
}

func CheckTypeUser(c *fiber.Ctx, typeUserID uint32) bool {
	return GetInfoUser(c).TypeUserID == typeUserID
}

func GetUserID(c *fiber.Ctx) uint32 {
	return GetInfoUser(c).UserID
}

func GetPermission(c *fiber.Ctx) []models.RolePermission {
	sess, _ := SessAuth.Get(c)
	if result, ok := sess.Get("permission").([]models.RolePermission); ok {
		return result
	}
	repo.OutPutDebugError("Failed to cast session data to permission")
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	repo.OutPutDebug("UpdateUser")

	userID := utilS.StringToUint32(c.Params("id"))

	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		repo.OutPutDebugError("User not found: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotFound, "User not found", nil)
	}

	var form models.UserWithoutPass
	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError("Body parsing failed: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Incorrect format", nil)
	}

	user.Name = form.Name
	user.Email = form.Email
	user.TypeUser = form.TypeUser
	user.Role = form.Role

	if err := utilS.UserRepo.Update(*user); err != nil {
		repo.OutPutDebugError("Failed to update user info: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Cannot update user info", nil)
	}

	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Update successfully",
		Data:    nil,
	})
}

func ApiGetUserByID(c *fiber.Ctx) error {
	userID := utilS.StringToUint32(c.Params("id"))
	repo.OutPutDebug("ApiGetUserByID")

	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}
func ApiUpdateUsers(c *fiber.Ctx) error {
	userID := utilS.StringToUint32(c.Params("id"))
	repo.OutPutDebug("ApiUpdateUsers")

	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	type Input struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		RoleID     string `json:"role_id"`
		TypeUserID string `json:"type_user_id"`
	}

	var input Input

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input data",
		})
	}

	user.Name = input.Name
	user.Email = input.Email
	user.RoleID = utilS.StringToUint32(input.RoleID)
	user.TypeUserID = utilS.StringToUint32(input.TypeUserID)

	if user.Name == "" || user.Email == "" || user.RoleID == 0 || user.TypeUserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Missing required fields",
		})
	}

	if err := utilS.UserRepo.Update(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
	})
}

func CreateUser(c *fiber.Ctx) error {
	repo.OutPutDebug("CreateUsers")
	return c.Render("pages/users/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}
func ApiCreateUsers(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateUsers")

	var form models.User
	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid request format", nil)
	}

	_, err := utilS.UserRepo.FindByEmail(form.Email)
	if err == nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Email already exists", nil)
	}

	var newUser models.User
	newUser.Email = form.Email
	newUser.Name = form.Name
	newUser.Password, newUser.Vault, _ = utilS.HashPassword(form.Password)
	newUser.TypeUserID = form.TypeUserID
	newUser.RoleID = form.RoleID
	newUser.IsActivated = false
	newUser.Token = ""
	newUser.TokenExpiry = time.Now()
	newUser.CreatedAt = time.Now()
	//newUser.CreatedBy = newUser.UserID
	if err := utilS.UserRepo.Create(newUser); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Create new user failed", nil)
	}

	//data := fiber.Map{
	//	"user": newUser,
	//}

	return c.Status(fiber.StatusCreated).JSON(&utilS.Response{
		Code:    fiber.StatusCreated,
		Message: "User created successfully with role and type user",
		Data:    nil,
	})
}

func ApiDeleteUser(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteUser")

	userID := utilS.StringToUint32(c.Params("id"))

	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		repo.OutPutDebugError("User not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	if err := utilS.UserRepo.Delete(*user); err != nil {
		repo.OutPutDebugError("Failed to delete user: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete user",
		})
	}

	repo.OutPutDebug("User deleted successfully")

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}

func GetUserInformation(c *fiber.Ctx) error {
	repo.OutPutDebug("GetUserInformation")
	u := GetInfoUser(c)

	user, err := utilS.UserRepo.FindByID(u.UserID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "User not found", nil)
	}

	data := fiber.Map{
		"Username": user.Name,
		"Email":    user.Email,
		"TypeUser": user.TypeUser.Name,
		"Vault":    user.TypeAccount,
		"Ctx":      c,
	}
	return c.Render("pages/users/information", data, "layouts/main")
}

func PutUserInformation(c *fiber.Ctx) error {
	repo.OutPutDebug("PutUserInformation")
	userLogin := GetInfoUser(c)

	user, _ := utilS.UserRepo.FindByID(userLogin.UserID)

	type Input struct {
		OddPassword     string `json:"odd_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	var input Input

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input data",
		})
	}

	if input.OddPassword != "" && !utilS.CheckPasswordHash(input.OddPassword, user.Vault, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Incorrect password",
		})
	}

	if len(input.NewPassword) < 8 && len(input.NewPassword) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Password requires 8 characters",
		})
	}

	if input.NewPassword != input.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Confirm password not match",
		})
	}

	if input.NewPassword != "" {
		user.Password, user.Vault, _ = utilS.HashPassword(input.NewPassword)
	}

	if err := utilS.UserRepo.Update(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "User updated successfully",
		Data:    nil,
	})
}

func ChangeEmail(c *fiber.Ctx) error {
	repo.OutPutDebug("ChangeEmail")
	userLogin := GetInfoUser(c)

	user, _ := utilS.UserRepo.FindByID(userLogin.UserID)

	type Input struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		OddPassword string `json:"odd_password"`
	}

	var input Input

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input data",
		})
	}

	if input.Email != user.Email {
		if user.Name == "" || user.Email == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Missing required fields",
			})
		}

		if _, err := utilS.UserRepo.FindByEmail(input.Email); err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Email already exists",
			})
		}

		token, _ := utilS.GenerateRandomString(128)
		user.Token = token
		user.TokenExpiry = time.Now().Add(time.Minute * 60)
		if err := utilS.UserRepo.Update(*user); err != nil {
			repo.OutPutDebugError(err.Error())
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot update new token", nil)
		}
		go func() {
			if err := utilS.SendEmailReqUpdateEmail(token, input.Email, user); err != nil {
			}
		}()

	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Request change email address successfully", nil)
}

func ApiConfirmEmail(c *fiber.Ctx) error {
	userID := utilS.StringToUint32(c.Params("id"))
	token := c.Params("token")
	newEmail := c.Params("email")

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

	user.Email = newEmail
	user.UpdatedAt = time.Now()

	if err := utilS.UserRepo.Update(*user); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot change email", nil)
	}

	return c.Redirect("/information")

	//return utilS.ResultResponse(c, fiber.StatusOK, "Forget password successfully. Please check your email for a new password.", nil)
}

func ApiChangePassword(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiChangePassword")
	userID := utilS.StringToUint32(c.Params("id"))
	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	type Input struct {
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	var input Input
	if err = c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid input data",
		})
	}

	if len(input.NewPassword) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Password requires 8 characters",
		})
	}

	if input.NewPassword != input.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Confirm password and New password not match",
		})
	}
	hashedPassword, vault, _ := utilS.HashPassword(input.NewPassword)
	user.Password = hashedPassword
	user.Vault = vault

	if err := utilS.UserRepo.Update(*user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    user,
		"success": true,
		"message": "Password updated successfully",
	})
}

func GetDetailUser(c *fiber.Ctx) error {
	repo.OutPutDebug("GetDetailUser")
	userIDParam := c.Params("id")
	userID := utilS.StringToUint32(userIDParam)
	user, err := utilS.UserRepo.FindByID(userID)
	if err != nil {
		repo.OutPutDebugError("User not found: " + err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.Render("pages/courses/student_detail", fiber.Map{
		"Ctx":    c,
		"UserID": userID,
		"User":   user,
	}, "layouts/main")
}

//
//func GetCourseLessons(c *fiber.Ctx) error {
//	repo.OutPutDebug("GetCourseLessons")
//
//	userIDParam := c.Params("id")
//	userID := utilS.StringToUint32(userIDParam)
//
//	courseIDParam := c.Params("course_id")
//	courseID := utilS.StringToUint32(courseIDParam)
//
//	user, err := utilS.UserRepo.FindByID(userID)
//	if err != nil {
//		repo.OutPutDebugError("User not found: " + err.Error())
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error": "User not found",
//		})
//	}
//
//	course, err := utilS.CourseRepo.FindByID(courseID)
//	if err != nil {
//		repo.OutPutDebugError("Course not found: " + err.Error())
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"error": "Course not found",
//		})
//	}
//
//	return c.Render("pages/courses/lessons_detail", fiber.Map{
//		"Ctx":      c,
//		"UserID":   userID,
//		"User":     user,
//		"Course":   course,
//		"CourseID": courseID,
//	}, "layouts/main")
//}
//
//func GetLessonAssignment(c *fiber.Ctx) error {
//	repo.OutPutDebug("GetLessonAssignment")
//
//	userIDParam := c.Params("id")
//	userID := utilS.StringToUint32(userIDParam)
//
//	lessonIDParam := c.Params("lesson_id")
//	lessonID := utilS.StringToUint32(lessonIDParam)
//
//	user, err := utilS.UserRepo.FindByID(userID)
//	if err != nil {
//		repo.OutPutDebugError("User not found: " + err.Error())
//		return utilS.ResultResponse(c, fiber.StatusNotFound, "User not found", nil)
//	}
//
//	lesson, err := utilS.LessonRepo.FindByID(lessonID)
//	if err != nil {
//		repo.OutPutDebugError("Lesson not found: " + err.Error())
//		return utilS.ResultResponse(c, fiber.StatusNotFound, "lesson not found", nil)
//	}
//
//	return c.Render("pages/courses/assignment_detail", fiber.Map{
//		"Ctx":      c,
//		"UserID":   userID,
//		"User":     user,
//		"Lesson":   lesson,
//		"LessonID": lessonID,
//	}, "layouts/main")
//}
//func GetLearningProcessByUser(c *fiber.Ctx) error {
//	repo.OutPutDebug("GetLearningProcessByUser")
//	userIDParam := c.Params("id")
//	userID := utilS.StringToUint32(userIDParam)
//
//	user, err := utilS.UserRepo.FindByID(userID)
//	if err != nil {
//		repo.OutPutDebugError("User not found: " + err.Error())
//		return utilS.ResultResponse(c, fiber.StatusNotFound, "User not found", nil)
//	}
//	return c.Render("pages/courses/learning_process", fiber.Map{
//		"Ctx":    c,
//		"UserID": userID,
//		"User":   user,
//	}, "layouts/main")
//}
