package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"strconv"
	"time"
)

func GetPrograms(c *fiber.Ctx) error {
	repo.OutPutDebug("GetPrograms")
	return c.Render("pages/programs/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}
func ApiGetPrograms(c *fiber.Ctx) error {
	repo.OutPutDebug("GetProgram")

	programs := utilS.ProgramRepo.FindAll()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Programs retrieved successfully",
		"data":    programs,
	})
}
func ApiGetProgramByID(c *fiber.Ctx) error {
	repo.OutPutDebug("GetProgramByID")
	programID := utilS.StringToUint32(c.Params("id"))

	program, err := utilS.ProgramRepo.FindByID(programID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "program not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"user":    program,
	})
}
func ApiCreateProgram(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateProgram")

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
	user := GetInfoUser(c)
	program := models.Program{
		Name:        request.Name,
		CreatedBy:   user.UserID,
		ProgramCode: utilS.GenerateCourseCode(3),
	}

	if err := utilS.ProgramRepo.Create(&program); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Failed to create program",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Program created successfully",
	})

}

func ApiDeleteProgram(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteProgram")

	programID := utilS.StringToUint32(c.Params("id"))
	program, err := utilS.ProgramRepo.FindByID(programID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Program not found",
		})
	}

	if err := utilS.ProgramRepo.Delete(program); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Failed to delete role",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Program deleted successfully",
	})

}

func ApiUpdateProgram(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateProgram")
	programID := utilS.StringToUint32(c.Params("id"))
	program, err := utilS.ProgramRepo.FindByID(programID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "program not found",
		})
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

	program.Name = request.Name
	program.UpdatedAt = time.Now()
	user := GetInfoUser(c)
	program.UpdatedBy = user.UserID

	if err := utilS.ProgramRepo.Update(*program); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Failed to update role",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Program updated successfully",
	})
}

func ApiGetPrerequisiteCourse(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetPrerequisiteCourse")
	programID := utilS.StringToUint32(c.Params("id"))
	courses, err := utilS.CourseRepo.FindByProgramID(programID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot get prerequisite courses", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Get prerequisite courses successfully", courses)
}

func GetStudentProgram(c *fiber.Ctx) error {
	repo.OutPutDebug("GetStudentProgram")
	return c.Render("pages/programs/student-programs", fiber.Map{"Ctx": c}, "layouts/main")
}

func GetStudentProgramDetails(c *fiber.Ctx) error {
	repo.OutPutDebug("GetStudentProgram")

	programID := utilS.StringToUint32(c.Params("id"))

	program, _ := utilS.ProgramRepo.FindByID(programID)

	courseUsers, _ := utilS.CourseRepo.FindByProgramID(programID)

	user := GetInfoUser(c)
	//courseUsers := utilS.CourseUserRepo.FindCourseByUserID(user.UserID, user.TypeUserID)

	type CourseDetails struct {
		CourseID        uint32
		CourseCode      string
		Title           string
		Description     string
		LessonCompleted uint32
		Lesson          uint32
		Status          bool
		Width           string
	}
	var courses []CourseDetails
	// foreach error
	for _, course := range courseUsers {
		courseTemp, _ := utilS.CourseRepo.FindByID(course.CourseID)
		countLesson, _ := utilS.LessonRepo.CountLessonInCourse(course.CourseID)
		countCompletedLesson, _ := utilS.UserLessonRepo.CountLessonCompletedInCourse(course.CourseID, user.UserID)
		width := float64(countCompletedLesson) / float64(countLesson) * 100
		var c CourseDetails
		c.CourseID = courseTemp.CourseID
		c.CourseCode = courseTemp.CourseCode
		c.Description = courseTemp.Description
		c.Title = courseTemp.Title
		c.Status = courseTemp.Status
		c.Lesson = countLesson
		c.LessonCompleted = countCompletedLesson
		c.Width = strconv.FormatFloat(width, 'f', 2, 64) + "%"
		courses = append(courses, c)
	}

	return c.Render("pages/programs/student-programs-details", fiber.Map{
		"Ctx":         c,
		"Courses":     courses,
		"ProgramName": program.Name,
	}, "layouts/main")
}

func ApiGetStudentProgram(c *fiber.Ctx) error {
	repo.OutPutDebug("GetStudentProgramDetails")

	user := GetInfoUser(c)

	userPrograms, err := utilS.UserProgramRepo.FindByUserID(user.UserID)

	type ProgramDetails struct {
		ProgramName string
		ProgramCode string
		ProgramID   uint32
		Course      uint32
		Status      string
	}

	var programs []ProgramDetails

	for _, userProgram := range userPrograms {
		count, _ := utilS.CourseRepo.CountCourseInProgram(userProgram.ProgramID)
		programs = append(programs, ProgramDetails{
			ProgramName: userProgram.Program.Name,
			ProgramID:   userProgram.ProgramID,
			ProgramCode: userProgram.Program.ProgramCode,
			Status:      "Learning",
			Course:      count,
		})
	}

	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot get student program details", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Get student program details successfully", fiber.Map{
		"UserProgram": programs,
		"Username":    user.Name,
	})
}

func slicesEqual(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func PostReqJoinProgram(c *fiber.Ctx) error {
	repo.OutPutDebug("PostReqJoinCourse")
	user := GetInfoUser(c)

	type RequestInput struct {
		ProgramCode string `json:"program_code"`
		ProgramID   string `json:"program_id"`
	}

	var (
		input           RequestInput
		userLessons     []*models.UserLesson
		userAssignments []*models.UserAssignment
		courseUsers     []models.CourseUser
	)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	var program *models.Program
	var err error
	if input.ProgramCode != "" {
		program, err = utilS.ProgramRepo.FindByProgramCode(input.ProgramCode)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "program not found",
			})
		}
	}

	if input.ProgramID != "" {
		program, err = utilS.ProgramRepo.FindByID(utilS.StringToUint32(input.ProgramID))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "program not found",
			})
		}
	}

	courseIDs, _ := utilS.CourseRepo.FindCourseIDByProgramID(program.ProgramID)
	existCourseID, _ := utilS.CourseUserRepo.FindCourseIDByProgramID(program.ProgramID, user.UserID)
	courses, _ := utilS.CourseRepo.FindInProgramID(courseIDs, existCourseID)
	lessonIDs, _ := utilS.LessonRepo.FindLessonIDByCourseIDs(courseIDs)
	assignments, _ := utilS.AssignmentRepo.FindInLessonID(lessonIDs)

	for _, course := range courses {
		courseUsers = append(courseUsers, models.CourseUser{
			CourseID: course.CourseID,
			UserID:   user.UserID,
		})
	}

	if err := utilS.CourseUserRepo.CreateBatch(courseUsers); err != nil {
		if slicesEqual(existCourseID, courseIDs) == false {
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"success": false,
				"message": "failed to create user course",
			})
		}
	}

	for _, lessonID := range lessonIDs {
		userLessons = append(userLessons, &models.UserLesson{
			LessonID: lessonID,
			UserID:   user.UserID,
		})
	}

	if err := utilS.UserLessonRepo.CreateAny(userLessons); err != nil {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"success": false,
			"message": "failed to create user lesson",
		})
	}

	for _, assignment := range assignments {
		userAssignments = append(userAssignments, &models.UserAssignment{
			UserID:       user.UserID,
			AssignmentID: assignment.AssignmentID,
		})
	}
	if err := utilS.UserAssignmentRepo.CreateAny(userAssignments); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to create user assignment",
		})
	}

	newReq := models.UserProgram{
		ProgramID: program.ProgramID,
		UserID:    user.UserID,
		//Status:   1,
	}

	_, err = utilS.UserProgramRepo.FindByUserIDAndProgramID(program.ProgramID, user.UserID)
	if err == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "program already exists",
		})
	}

	err = utilS.UserProgramRepo.Create(newReq)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "create request failed", nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Create request successfully",
		"data":    nil,
	})
}

func GetProgramDetails(c *fiber.Ctx) error {
	repo.OutPutDebug("GetStudentProgram")
	return c.Render("pages/programs/details", fiber.Map{"Ctx": c}, "layouts/main")
}

func ApiGetCourseInProgram(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourseInProgram")

	programID := utilS.StringToUint32(c.Params("id"))

	courses, err := utilS.CourseRepo.FindByProgramID(programID)

	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot get courses", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusAccepted, "Cannot get courses", courses)
}

func GetSearchPrograms(c *fiber.Ctx) error {
	id := c.Params("id")
	user := GetInfoUser(c)
	_, err := utilS.UserProgramRepo.FindByUserIDAndProgramID(utilS.StringToUint32(id), user.UserID)

	program, _ := utilS.ProgramRepo.FindByID(utilS.StringToUint32(id))

	location := "/student-programs/" + id + "/details"

	courseUsers, _ := utilS.CourseRepo.FindByProgramID(utilS.StringToUint32(id))

	if err != nil {
		type CourseDetails struct {
			CourseID        uint32
			CourseCode      string
			Title           string
			Description     string
			LessonCompleted uint32
			Lesson          uint32
			Status          bool
			Width           string
		}
		var courses []CourseDetails
		// foreach error
		for _, course := range courseUsers {
			courseTemp, _ := utilS.CourseRepo.FindByID(course.CourseID)
			countLesson, _ := utilS.LessonRepo.CountLessonInCourse(course.CourseID)
			countCompletedLesson, _ := utilS.UserLessonRepo.CountLessonCompletedInCourse(course.CourseID, user.UserID)
			width := float64(countCompletedLesson) / float64(countLesson) * 100
			var c CourseDetails
			c.CourseID = courseTemp.CourseID
			c.CourseCode = courseTemp.CourseCode
			c.Description = courseTemp.Description
			c.Title = courseTemp.Title
			c.Status = courseTemp.Status
			c.Lesson = countLesson
			c.LessonCompleted = countCompletedLesson
			c.Width = strconv.FormatFloat(width, 'f', 2, 64) + "%"
			courses = append(courses, c)
		}

		return c.Render("pages/programs/search-programs", fiber.Map{
			"Ctx":         c,
			"Courses":     courses,
			"ProgramName": program.Name,
		}, "layouts/main")
	} else {
		return c.Redirect(location)
	}
}
