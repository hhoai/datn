package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"time"
)

func GetCourseCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLevel")
	return c.Render("pages/course_category/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}

func ApiGetCourseCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourseCategory")

	courseCategories := utilS.CourseCategoryRepo.FindAll()

	return utilS.ResultResponse(c, fiber.StatusOK, "Category retrieved successfully", courseCategories)
}

func ApiGetCourseCategoryByID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourseCategoryByID")
	courseCategoryID := utilS.StringToUint32(c.Params("id"))

	courseCategory, err := utilS.CourseCategoryRepo.FindByID(courseCategoryID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Category not found", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Category retrieved successfully", courseCategory)
}
func ApiCreateCourseCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateCourseCategory")

	type Request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid format", nil)
	}
	user := GetInfoUser(c)
	courseCategory := models.CourseCategory{
		Name:        request.Name,
		Description: request.Description,
		CreatedBy:   user.UserID,
	}

	if err := utilS.CourseCategoryRepo.Create(&courseCategory); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusCreated, "Skills create successfully", courseCategory)

}

func ApiDeleteCourseCategory(c *fiber.Ctx) error {
	repo.OutPutDebug(" ApiDeleteCourseCategory")

	courseCategoryID := utilS.StringToUint32(c.Params("id"))

	courseCategory, err := utilS.CourseCategoryRepo.FindByID(courseCategoryID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Category not found", nil)
	}
	if err := utilS.CourseCategoryRepo.Delete(courseCategory); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to delete", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Category deleted successful", nil)

}
func ApiUpdateCourseCategory(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateCourseCategory")
	courseCategoryID := utilS.StringToUint32(c.Params("id"))

	courseCategory, err := utilS.CourseCategoryRepo.FindByID(courseCategoryID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Category not found", nil)
	}
	type Request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid request",
		})
	}

	courseCategory.Name = request.Name
	courseCategory.Description = request.Description
	courseCategory.UpdatedAt = time.Now()
	user := GetInfoUser(c)
	courseCategory.UpdatedBy = user.UserID

	if err := utilS.CourseCategoryRepo.Update(*courseCategory); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to update", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Category updated ok ", courseCategory)
}
