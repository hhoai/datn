package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/structs"
	"lms/utilS"
	"time"
)

func GetCourses(c *fiber.Ctx) error {
	repo.OutPutDebug("GetCourses")
	return c.Render("pages/courses/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}
func ApiGetCourses(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourses")

	courses := utilS.CourseRepo.FindAll()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Users retrieved successfully",
		"data":    courses,
	})
}

func ApiCreateCourse(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateCourse")

	var form structs.FormCreateCourse
	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid request format", nil)
	}

	startTime, err := time.Parse("2006-01-02", form.StartTime)
	if err != nil {
		repo.OutPutDebugError("Invalid StartTime format")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid StartTime format. Expected format: YYYY-MM-DD", nil)
	}

	// Parse EndTime
	endTime, err := time.Parse("2006-01-02", form.EndTime)
	if err != nil {
		repo.OutPutDebugError("Invalid EndTime format")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid EndTime format. Expected format: YYYY-MM-DD", nil)
	}

	if endTime.Before(startTime) {
		repo.OutPutDebug("Invalid EndTime format")
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Start time must not be less than end time", nil)
	}

	program, err := utilS.ProgramRepo.FindByID(form.ProgramID)
	if err != nil {
		repo.OutPutDebugError("Invalid ProgramID")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid ProgramID", nil)
	}

	level, err := utilS.LevelRepo.FindByID(form.LevelID)
	if err != nil {
		repo.OutPutDebugError("Invalid LevelID")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid LevelID", nil)
	}

	courseCategory, err := utilS.CourseCategoryRepo.FindByID(form.CourseCategoriesID)
	if err != nil {
		repo.OutPutDebugError("Invalid CourseCategoryID")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid CourseCategoryID", nil)
	}

	var user models.UserWithoutPass
	user = GetInfoUser(c)
	newCourse := models.Course{
		Title:                form.Title,
		Description:          form.Description,
		LevelID:              form.LevelID,
		ProgramID:            form.ProgramID,
		CourseCategoriesID:   form.CourseCategoriesID,
		Image:                form.Image,
		Amount:               form.Amount,
		StartTime:            startTime,
		EndTime:              endTime,
		CreatedBy:            user.UserID,
		UpdatedBy:            user.UserID,
		Status:               false,
		CourseCode:           utilS.GenerateCourseCode(3),
		PrerequisiteCourseID: form.PrerequisiteCourseID,
	}

	if err = utilS.CourseRepo.Create(newCourse); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create course", nil)
	}

	newCourse.Program = *program
	newCourse.Level = *level
	newCourse.CourseCategory = *courseCategory

	data := fiber.Map{
		"course": newCourse,
	}

	return c.Status(fiber.StatusOK).JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Course created successfully",
		Data:    data,
	})
}
func ApiGetCourseById(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourseById")
	courseID := utilS.StringToUint32(c.Params("id"))
	course, err := utilS.CourseRepo.FindByID(courseID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "course not found",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"course":  course,
	})
}
func ApiUpdateCourse(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateCourse")

	courseID := utilS.StringToUint32(c.Params("id"))

	course, err := utilS.CourseRepo.FindByID(courseID)
	if err != nil {
		repo.OutPutDebugError("Course not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Course not found",
		})
	}

	var form structs.FormCreateCourse
	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError("Error parsing request body: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid request format", nil)
	}

	repo.OutPutDebug(fmt.Sprintf("Received form: %+v", form))

	startTime, err := time.Parse("2006-01-02", form.StartTime)
	if err != nil {
		repo.OutPutDebugError("Invalid StartTime format")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid StartTime format. Expected format: YYYY-MM-DD", nil)
	}

	endTime, err := time.Parse("2006-01-02", form.EndTime)
	if err != nil {
		repo.OutPutDebugError("Invalid EndTime format")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid EndTime format. Expected format: YYYY-MM-DD", nil)
	}
	if endTime.Before(startTime) {
		repo.OutPutDebug("Invalid EndTime format")
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Start time must not be less than end time", nil)
	}
	program, err := utilS.ProgramRepo.FindByID(form.ProgramID)
	if err != nil {
		repo.OutPutDebugError("Invalid ProgramID")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid ProgramID", nil)
	}

	level, err := utilS.LevelRepo.FindByID(form.LevelID)
	if err != nil {
		repo.OutPutDebugError("Invalid LevelID")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid LevelID", nil)
	}

	courseCategory, err := utilS.CourseCategoryRepo.FindByID(form.CourseCategoriesID)
	if err != nil {
		repo.OutPutDebugError("Invalid CourseCategoryID")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid CourseCategoryID", nil)
	}
	if form.Status != false && form.Status != true {
		repo.OutPutDebugError("Invalid status value")
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid status value. Allowed values: 0 or 1", nil)
	}

	//if course.Status == false && form.Status == true {
	//	users, _ := utilS.UserRepo.FindByTypeUserID(4)
	//	utilS.SendEmailNewCourse(course.CourseCode, users)
	//}

	user := GetInfoUser(c)
	course.Title = form.Title
	course.Description = form.Description
	course.LevelID = form.LevelID
	course.ProgramID = form.ProgramID
	course.CourseCategoriesID = form.CourseCategoriesID
	course.Image = form.Image
	course.Amount = form.Amount
	course.Status = form.Status
	course.StartTime = startTime
	course.EndTime = endTime
	course.UpdatedBy = user.UserID
	course.PrerequisiteCourseID = form.PrerequisiteCourseID

	if err := utilS.CourseRepo.Update(*course); err != nil {
		repo.OutPutDebugError("Failed to update course: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to update course", nil)
	}

	course.Program = *program
	course.Level = *level
	course.CourseCategory = *courseCategory

	data := fiber.Map{
		"course": course,
	}

	return c.Status(fiber.StatusOK).JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Course updated successfully",
		Data:    data,
	})
}

func ApiDeleteCourse(c *fiber.Ctx) error {
	repo.OutPutDebug("DeleteCourse")
	courseID := utilS.StringToUint32(c.Params("id"))

	course, err := utilS.CourseRepo.FindByID(courseID)
	if err != nil {
		repo.OutPutDebugError("course not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "course not found",
		})
	}

	if err := utilS.CourseRepo.Delete(*course); err != nil {
		repo.OutPutDebugError("Failed to delete course: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete course",
		})
	}

	repo.OutPutDebug("course deleted successfully")

	return c.JSON(fiber.Map{
		"success": true,
		"message": "course deleted successfully",
	})
}
func GetCourseDetails(c *fiber.Ctx) error {
	repo.OutPutDebug("GetCourseDetails")
	p := c.Params("id")

	courseID := utilS.StringToUint32(p)

	course, err := utilS.CourseRepo.FindByID(courseID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
		return err
	}

	lessons, _ := utilS.LessonRepo.FindByCourseID(courseID)

	lessonCategories := utilS.LessonCategoryRepo.FindAll()

	level := utilS.LevelRepo.FindAll()

	data := fiber.Map{
		"CourseID":         courseID,
		"Title":            course.Title,
		"Description":      course.Description,
		"Ctx":              c,
		"Lessons":          lessons,
		"Level":            level,
		"LessonCategories": lessonCategories,
	}

	return c.Render("pages/courses/details", data, "layouts/main")
}

func GetStudentAssignmentCourseTopicID(c *fiber.Ctx) error {
	return c.Render("pages/topics/topic-detail", fiber.Map{"Ctx": c}, "layouts/main")
}

func DeleteMultipleCourses(c *fiber.Ctx) error {
	type RequestBody struct {
		CourseIDs []uint32 `json:"course_id"`
	}

	var reqBody RequestBody

	if err := c.BodyParser(&reqBody); err != nil {
		repo.OutPutDebugError("Failed to parse body: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Incorrect format", nil)
	}

	if len(reqBody.CourseIDs) == 0 {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Courses not found", nil)
	}

	err := utilS.CourseRepo.DeleteMultiple(reqBody.CourseIDs)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}

	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Course deleted successfully",
		Data:    nil,
	})
}

func ApiGetCompletedLessons(c *fiber.Ctx) error {
	p := c.Params("id")
	userID := utilS.StringToUint32(p)

	completedLessons, err := utilS.UserLessonRepo.FindCompletedLesson(userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to fetch completed lessons", nil)
	}

	type CompletedLessonFormat struct {
		CourseID     uint32 `json:"course_id"`
		CourseName   string `json:"course_name"`
		LessonID     uint32 `json:"lesson_id"`
		LessonName   string `json:"lesson_name"`
		TimelineDate string `json:"timeline_date"`
		SubmitAt     string `json:"submit_at"`
	}

	var completedCourses []CompletedLessonFormat

	for _, lesson := range completedLessons {
		completedCourses = append(completedCourses, CompletedLessonFormat{
			CourseID:     lesson.CourseID,
			CourseName:   lesson.CourseName,
			LessonID:     lesson.LessonID,
			LessonName:   lesson.LessonName,
			TimelineDate: lesson.TimelineDate.Format("2006-01-02"),
			SubmitAt:     lesson.SubmitAt.Format("2006-01-02 15:04:05.000"),
		})
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Get Completed Lesson successfully", completedCourses)
}

//func GetLessonByCourseID(c *fiber.Ctx) error {
//	repo.OutPutDebug("GetLessonByCourseID")
//	return c.Render("pages/users/index", fiber.Map{
//		"Ctx": c,
//	}, "layouts/main")
//}

func ApiGetCompletedAssignment(c *fiber.Ctx) error {
	p := c.Params("id")
	userID := utilS.StringToUint32(p)

	completedAssignments, err := utilS.UserAssignmentRepo.FindCompletedAssignment(userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to fetch completed assignments", nil)
	}
	//struct data response

	type CompletedAssignmentFormat struct {
		CourseID       uint32 `json:"course_id"`
		CourseName     string `json:"course_name"`
		LessonID       uint32 `json:"lesson_id"`
		LessonName     string `json:"lesson_name"`
		AssignmentID   uint32 `json:"assignment_id"`
		AssignmentName string `json:"assignment_name"`
		TimelineDate   string `json:"timeline_date"`
		SubmitAt       string `json:"submit_at"`
	}

	var formattedAssignments []CompletedAssignmentFormat
	for _, assignment := range completedAssignments {
		formattedAssignments = append(formattedAssignments, CompletedAssignmentFormat{
			CourseID:       assignment.CourseID,
			CourseName:     assignment.CourseName,
			LessonID:       assignment.LessonID,
			LessonName:     assignment.LessonName,
			AssignmentID:   assignment.AssignmentID,
			AssignmentName: assignment.AssignmentName,
			TimelineDate:   assignment.TimelineDate.Format("2006-01-02"),
			SubmitAt:       assignment.SubmitAt.Format("2006-01-02 15:04:05"),
		})
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Get completed assignments successfully", formattedAssignments)
}
func ApiGetCompletedCourse(c *fiber.Ctx) error {
	p := c.Params("id")
	userID := utilS.StringToUint32(p)

	completedCourses, err := utilS.CourseUserRepo.FindCompletedCourse(userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to fetch completed courses", nil)
	}
	//struct data response

	type CompletedCourseFormat struct {
		CourseID     uint32 `json:"course_id"`
		CourseName   string `json:"course_name"`
		TimelineDate string `json:"timeline_date"`
		SubmitAt     string `json:"submit_at"`
	}

	var formattedCourses []CompletedCourseFormat
	for _, completedCourse := range completedCourses {
		formattedCourses = append(formattedCourses, CompletedCourseFormat{
			CourseID:     completedCourse.CourseID,
			CourseName:   completedCourse.Course.Title,
			TimelineDate: completedCourse.CompletedAt.Format("2006-01-02"),
			SubmitAt:     completedCourse.CompletedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Get Completed Courses successfully", formattedCourses)
}
func ApiGetStatisticsCourse(c *fiber.Ctx) error {
	repo.OutPutDebugError("ApiGetStatisticsCourse")
	userIDParam := c.Params("id")
	userID := utilS.StringToUint32(userIDParam)

	courseIDParam := c.Params("course_id")
	courseID := utilS.StringToUint32(courseIDParam)

	totalLessons, err := utilS.LessonRepo.CountLessonInCourse(courseID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to count lessons in course", nil)
	}

	completedLessons, err := utilS.UserLessonRepo.CountLessonCompletedInCourse(courseID, userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to count complete lessons in course", nil)
	}
	data := fiber.Map{
		"total_lessons":     totalLessons,
		"completed_lessons": completedLessons,
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Get Statistics successfully", data)

}

func ApiGetStatisticsLesson(c *fiber.Ctx) error {
	repo.OutPutDebugError("ApiGetStatisticsLesson")

	userIDParam := c.Params("id")
	userID := utilS.StringToUint32(userIDParam)

	lessonIDParam := c.Params("lesson_id")
	lessonID := utilS.StringToUint32(lessonIDParam)

	totalAssignments, err := utilS.AssignmentRepo.CountAssignmentInLesson(lessonID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to count assignment in lesson ", nil)
	}

	completedAssignments, err := utilS.UserAssignmentRepo.CountAssignmentCompletedInLesson(lessonID, userID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to count complete assignment in lesson", nil)
	}
	data := fiber.Map{
		"totalAssignments":     totalAssignments,
		"completedAssignments": completedAssignments,
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Get Statistics successfully", data)

}
