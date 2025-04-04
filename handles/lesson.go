package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/structs"
	"lms/utilS"
	"time"
)

func ApiGetLessonByCourseID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourses")

	courseID := utilS.StringToUint32(c.Params("id"))

	lessons, _ := utilS.LessonRepo.FindByCourseID(courseID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Courses retrieved successfully",
		"data":    lessons,
	})
}

func ApiGetLessons(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourses")

	lessons := utilS.LessonRepo.FindAll()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Courses retrieved successfully",
		"data":    lessons,
	})
}

func GetLessonDetails(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLessonDetails")
	p := c.Params("id")

	lessonID := utilS.StringToUint32(p)

	lesson, err := utilS.LessonRepo.FindByID(lessonID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
		return err
	}

	lessonCategories := utilS.LessonCategoryRepo.FindAll()
	typeAssignments := utilS.TypeAssignmentRepo.FindAll()

	data := fiber.Map{
		"LessonID":         lessonID,
		"Title":            lesson.Title,
		"Level":            lesson.Level,
		"Ctx":              c,
		"TypeAssignments":  typeAssignments,
		"LessonCategories": lessonCategories,
		"LessonCategoryID": lesson.LessonCategoryID,
	}

	return c.Render("pages/lesson/index", data, "layouts/main")
}

func GetPosts(c *fiber.Ctx) error {
	repo.OutPutDebug("GetPosts")
	p := c.Params("id")

	lessonID := utilS.StringToUint32(p)

	lesson, err := utilS.LessonRepo.FindByID(lessonID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
		return err
	}

	lessonCategories := utilS.LessonCategoryRepo.FindAll()
	typeAssignments := utilS.TypeAssignmentRepo.FindAll()

	data := fiber.Map{
		"LessonID":         lessonID,
		"Title":            lesson.Title,
		"Level":            lesson.Level,
		"Ctx":              c,
		"TypeAssignments":  typeAssignments,
		"LessonCategories": lessonCategories,
		"LessonCategoryID": lesson.LessonCategoryID,
	}

	return c.Render("pages/lesson/posts", data, "layouts/main")
}

func CreateLesson(c *fiber.Ctx) error {
	repo.OutPutDebug("CreateLesson")
	var (
		newLesson   models.Lesson
		user        models.UserWithoutPass
		form        structs.CreateLessonForm
		userLessons []*models.UserLesson
	)

	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError("Failed to parse body: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Incorrect format", nil)
	}
	user = GetInfoUser(c)

	newLesson.Title = form.Title
	newLesson.LevelID = utilS.StringToUint32(form.LevelID)
	newLesson.CourseID = utilS.StringToUint32(form.CourseID)
	newLesson.LessonCategoryID = utilS.StringToUint32(form.LessonCategoriesID)
	newLesson.CreatedAt = time.Now()
	newLesson.UpdatedAt = time.Now()
	newLesson.CreatedBy = user.UserID
	newLesson.UpdatedBy = user.UserID

	if err := utilS.LessonRepo.Create(newLesson); err != nil {
		repo.OutPutDebugError("Error creating lessons: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Error creating lessons", nil)
	}
	lesson, _ := utilS.LessonRepo.GetLastLesson()
	courseUsers := utilS.CourseUserRepo.FindUserByCourseID(utilS.StringToUint32(form.CourseID), 4)

	for _, courseUser := range courseUsers {
		userLessons = append(userLessons, &models.UserLesson{
			UserID:   courseUser.UserID,
			LessonID: lesson.LessonID,
		})
	}
	if userLessons != nil {
		if err := utilS.UserLessonRepo.CreateAny(userLessons); err != nil {
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Error creating lessons", nil)
		}
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Lesson created successfully", nil)
}

func UpdateLesson(c *fiber.Ctx) error {
	repo.OutPutDebug("UpdateLesson")
	var user models.UserWithoutPass
	user = GetInfoUser(c)

	lessonID := utilS.StringToUint32(c.Params("id"))

	lesson, _ := utilS.LessonRepo.FindByID(lessonID)

	var form structs.CreateLessonForm

	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNoContent, "Incorrect format", nil)
	}

	lesson.Title = form.Title
	lesson.LevelID = utilS.StringToUint32(form.LevelID)
	lesson.LessonCategoryID = utilS.StringToUint32(form.LessonCategoriesID)
	lesson.UpdatedBy = user.UserID
	lesson.UpdatedAt = time.Now()

	if err := utilS.LessonRepo.Update(*lesson); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update lessons", nil)
	}

	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Update successfully",
		Data:    nil,
	})
}

func GetLessonByID(c *fiber.Ctx) error {
	lessonID := utilS.StringToUint32(c.Params("id"))

	lesson, err := utilS.LessonRepo.FindByID(lessonID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Get lesson successfully",
		Data:    fiber.Map{"lesson": lesson},
	})
}

func DeleteLesson(c *fiber.Ctx) error {
	repo.OutPutDebug("DeleteLesson")
	lessonID := utilS.StringToUint32(c.Params("id"))

	err := utilS.LessonRepo.DeleteByID(lessonID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Lesson deleted successfully",
		Data:    nil,
	})
}
