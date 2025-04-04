package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"time"
)

func GetLessonCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLevel")
	return c.Render("pages/lesson_category/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}

func ApiGetLessonCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetLessonCategory")

	lessonCategory := utilS.LessonCategoryRepo.FindAll()

	return utilS.ResultResponse(c, fiber.StatusOK, "Category retrieved successfully", lessonCategory)
}

func ApiGetLessonCategoryByID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetLessonCategoryByID")
	lessonCategoryID := utilS.StringToUint32(c.Params("id"))

	lessonCategory, err := utilS.LessonCategoryRepo.FindByID(lessonCategoryID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Category not found", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Category retrieved successfully", lessonCategory)
}
func ApiCreateLessonCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateLessonCategory")

	type Request struct {
		Name string `json:"name"`
	}
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid format", nil)
	}
	user := GetInfoUser(c)
	lessonCategory := models.LessonCategory{
		Name:      request.Name,
		CreatedBy: user.UserID,
	}

	if err := utilS.LessonCategoryRepo.Create(&lessonCategory); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusCreated, "Skills create successfully", lessonCategory)

}

func ApiDeleteLessonCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteLessonCategory")

	lessonCategoryID := utilS.StringToUint32(c.Params("id"))

	lessonCategory, err := utilS.LessonCategoryRepo.FindByID(lessonCategoryID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Category not found", nil)
	}
	if err := utilS.LessonCategoryRepo.Delete(lessonCategory); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to delete", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Category deleted successful", nil)

}
func ApiUpdateLessonCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateCourseCategory")
	lessonCategoryID := utilS.StringToUint32(c.Params("id"))

	lessonCategory, err := utilS.LessonCategoryRepo.FindByID(lessonCategoryID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Category not found", nil)
	}
	type Request struct {
		Name string `json:"name"`
	}
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid request",
		})
	}

	lessonCategory.Name = request.Name
	lessonCategory.UpdatedAt = time.Now()
	user := GetInfoUser(c)
	lessonCategory.UpdatedBy = user.UserID

	if err := utilS.LessonCategoryRepo.Update(*lessonCategory); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to update", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Category updated ok ", lessonCategory)
}
