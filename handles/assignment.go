package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"lms/models"
	"lms/repo"
	"lms/structs"
	"lms/utilS"
	"time"
)

const EmailSMTP = "noreply.bitcare@gmail.com"

func ApiGetAssignmentByLessonID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourses")
	lessonID := utilS.StringToUint32(c.Params("id"))

	assignments, _ := utilS.AssignmentRepo.FindByLessonID(lessonID)

	lessons, _ := utilS.LessonRepo.FindByID(lessonID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "assignments retrieved successfully",
		"data": fiber.Map{
			"assignments": assignments,
			"lessons":     lessons,
		},
	})
}

func ApiGetAssignments(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourses")

	assignments := utilS.AssignmentRepo.FindAll()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Courses retrieved successfully",
		"data":    assignments,
	})
}

func GetAssignmentDetails(c *fiber.Ctx) error {
	repo.OutPutDebug("GetAssignmentDetails")
	//p := c.Params("id")
	//
	//assignmentID := utilS.StringToUint32(p)
	//
	//assignment, err := utilS.AssignmentRepo.FindByID(assignmentID)
	//if err != nil {
	//	repo.OutPutDebugError(err.Error())
	//	return err
	//}
	//
	//file, err := utilS.FileAssignmentRepo.FindByAssignmentID(assignmentID)
	//
	data := fiber.Map{
		//"AssignmentID":    assignmentID,
		//"AssignmentTitle": assignment.Title,
		//"AssignmentBody":  assignment.Body,
		//"DueDate":         assignment.DueDate.Format("2006-01-02T15:04"),
		//"Status":          assignment.Status,
		//"Score":           assignment.Score,
		//"Files":           file,
		"Ctx": c,
	}

	return c.Render("pages/assignment/detail", data, "layouts/main")
}

// send email
func SendEmailToStudents(to []string, subject, body string) error {
	for _, email := range to {
		go func(email string, subject, body string) {
			mailer := gomail.NewMessage()
			mailer.SetHeader("From", EmailSMTP)
			mailer.SetHeader("To", email)
			mailer.SetHeader("Subject", subject)
			mailer.SetBody("text/html", body)

			dialer := gomail.NewDialer("smtp.gmail.com", 587, EmailSMTP, "dlmd wcie mrnv ectj")
			dialer.DialAndSend(mailer)
		}(email, subject, body)

	}
	return nil
}

func CreateAssignment(c *fiber.Ctx) error {
	repo.OutPutDebug("CreateAssignment")

	var (
		newAssignment   models.Assignment
		user            models.UserWithoutPass
		form            structs.CreateAssignmentForm
		userAssignments []*models.UserAssignment
	)
	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError("Failed to parse body: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Incorrect format", nil)
	}

	user = GetInfoUser(c)

	newAssignment.Title = form.Title
	newAssignment.Body = form.Description
	newAssignment.LessonID = utilS.StringToUint32(form.LessonID)
	newAssignment.TypeAssignmentID = utilS.StringToUint32(form.TypeAssignmentID)
	newAssignment.DueDate = time.Now()
	newAssignment.CreatedAt = time.Now()
	newAssignment.UpdatedAt = time.Now()
	newAssignment.CreatedBy = user.UserID
	newAssignment.UpdatedBy = user.UserID

	if err := utilS.AssignmentRepo.Create(newAssignment); err != nil {
		repo.OutPutDebugError("Error creating lessons: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Error creating assignment", nil)
	}

	lesson, _ := utilS.LessonRepo.FindByID(newAssignment.LessonID)

	courseUsers := utilS.CourseUserRepo.FindUserByCourseID(lesson.CourseID, 4)

	var studentEmail []string
	for _, student := range courseUsers {
		if student.User.TypeUserID == 4 {
			studentEmail = append(studentEmail, student.User.Email)
		}
	}

	title := "Thông báo về bài tập"
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
			<p>Giáo viên vừa đăng tải một bài tập mới trên hệ thống.</p>
		</div>
	</body>
	</html>
`

	if studentEmail != nil {
		if err := SendEmailToStudents(studentEmail, title, body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
		}
		assignment, _ := utilS.AssignmentRepo.GetLastAssignment()
		userLesson, _ := utilS.UserLessonRepo.FindByLessonID(utilS.StringToUint32(form.LessonID))
		for _, userLesson := range userLesson {
			userAssignments = append(userAssignments, &models.UserAssignment{
				UserID:       userLesson.UserID,
				AssignmentID: assignment.AssignmentID,
			})
		}
		if err := utilS.UserAssignmentRepo.CreateAny(userAssignments); err != nil {
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Error creating assignment", nil)
		}
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Assignment created successfully", nil)
}

func UpdateAssignment(c *fiber.Ctx) error {
	repo.OutPutDebug("UpdateAssignment")
	var user models.UserWithoutPass
	user = GetInfoUser(c)

	assignmentID := utilS.StringToUint32(c.Params("id"))
	assignment, _ := utilS.AssignmentRepo.FindByID(assignmentID)
	var form structs.UpdateAssignmentForm

	//if err := c.BodyParser(&form); err != nil {
	//	repo.OutPutDebugError(err.Error())
	//	return utilS.ResultResponse(c, fiber.StatusNoContent, "Incorrect format", nil)
	//}

	form.Title = c.FormValue("title")
	form.Body = c.FormValue("body")
	form.Score = c.FormValue("score")
	form.DueDate = c.FormValue("due_date")
	assignment.Title = form.Title
	assignment.Body = form.Body
	assignment.DueDate = utilS.StringToTime(form.DueDate)
	assignment.Score = utilS.StringToUint32(form.Score)
	assignment.TypeAssignmentID = utilS.StringToUint32(c.FormValue("type_assignment_id"))
	assignment.UpdatedBy = user.UserID
	assignment.UpdatedAt = time.Now()
	assignment.CreatedBy = user.UserID
	//create folder assignments
	filepath := "./documents/assignments/" + utilS.Uint32ToString(assignment.LessonID)
	if err := utilS.CreateFolder(filepath); err != nil {
		repo.OutPutDebugError(err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		//return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update file", nil)
		repo.OutPutDebugError(err.Error())
	} else {
		err = c.SaveFile(file, filepath+"/"+file.Filename)
		if err != nil {
			//return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update file", nil)
			repo.OutPutDebugError(err.Error())
		}

		var fileAssignment models.FileAssignment
		fileAssignment.FileName = file.Filename
		fileAssignment.AssignmentID = assignmentID
		fileAssignment.CreatedBy = assignment.CreatedBy
		if err := utilS.FileAssignmentRepo.Create(fileAssignment); err != nil {
			repo.OutPutDebugError(err.Error())
			return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update file", nil)
		}
	}
	if err := utilS.AssignmentRepo.Update(*assignment); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update assignment", nil)
	}
	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Update successfully",
		Data:    nil,
	})
}

func GetAssignmentByID(c *fiber.Ctx) error {
	assignmentID := utilS.StringToUint32(c.Params("id"))

	assignment, err := utilS.AssignmentRepo.FindByID(assignmentID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Get lesson successfully",
		Data:    fiber.Map{"assignment": assignment},
	})
}

func DeleteAssignment(c *fiber.Ctx) error {
	repo.OutPutDebug("DeleteAssignment")
	assignmentID := utilS.StringToUint32(c.Params("id"))

	err := utilS.AssignmentRepo.DeleteByID(assignmentID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Assignment deleted successfully",
		Data:    nil,
	})
}

func ApiGetStudentAssignmentDetails(c *fiber.Ctx) error {
	assignmentID := utilS.StringToUint32(c.Params("id"))
	assignment, err := utilS.AssignmentRepo.FindByID(assignmentID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}

	fileAssignment, _ := utilS.FileAssignmentRepo.FindByAssignmentIDAndUserID(assignmentID, assignment.CreatedBy)

	user := GetInfoUser(c)

	fileStudentAssignment, err := utilS.FileAssignmentRepo.FindByAssignmentIDAndUserID(assignmentID, user.UserID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}

	lesson, _ := utilS.LessonRepo.FindByID(assignment.LessonID)
	course, _ := utilS.CourseRepo.FindByID(lesson.CourseID)

	courseStatus := course.Status

	userAssignment, _ := utilS.UserAssignmentRepo.FindByUserIDAndAssignmentID(assignmentID, user.UserID)

	return utilS.ResultResponse(c, fiber.StatusOK, "Student Assignment Details", fiber.Map{
		"Assignment":            assignment,
		"FileStudentAssignment": fileStudentAssignment,
		"FileAssignment":        fileAssignment,
		"Lesson":                lesson,
		"UserAssignment":        userAssignment,
		"CourseTitle":           course.Title,
		"CourseStatus":          courseStatus,
	})
}

func GetStudentAssignmentDetails(c *fiber.Ctx) error {

	return c.Render("pages/assignment/student-assignment", fiber.Map{"Ctx": c}, "layouts/main")
}

func GetStudentAssignmentTopicID(c *fiber.Ctx) error {
	return c.Render("pages/topics/student-detail", fiber.Map{"Ctx": c}, "layouts/main")
}

func ApiSubmitAssignment(c *fiber.Ctx) error {
	assignmentID := c.FormValue("assignment_id")
	user := GetInfoUser(c)
	//create folder assignments
	filepath := "./documents/student/" + utilS.Uint32ToString(user.UserID) + "/" + assignmentID
	if err := utilS.CreateFolder(filepath); err != nil {
		repo.OutPutDebugError(err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		//return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update file", nil)
	} else {
		err = c.SaveFile(file, filepath+"/"+file.Filename)
		if err != nil {
			//return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update file", nil)
		}

		var fileAssignment models.FileAssignment
		fileAssignment.FileName = file.Filename
		fileAssignment.AssignmentID = utilS.StringToUint32(assignmentID)
		fileAssignment.CreatedBy = user.UserID

		if err := utilS.FileAssignmentRepo.Create(fileAssignment); err != nil {
			repo.OutPutDebugError(err.Error())
			return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update file", nil)
		}
	}

	userAssignment, _ := utilS.UserAssignmentRepo.FindByUserIDAndAssignmentID(utilS.StringToUint32(assignmentID), user.UserID)
	userAssignment.Status = true
	userAssignment.CompletedAt = time.Now()

	if err := utilS.UserAssignmentRepo.Update(*userAssignment); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update user assignment", nil)
	}

	fileAssignments, _ := utilS.FileAssignmentRepo.FindByAssignmentIDAndUserID(utilS.StringToUint32(assignmentID), user.UserID)
	userAssignment, _ = utilS.UserAssignmentRepo.FindByUserIDAndAssignmentID(utilS.StringToUint32(assignmentID), user.UserID)
	lesson, _ := utilS.LessonRepo.FindByID(userAssignment.Assignment.LessonID)
	countAssignment, _ := utilS.AssignmentRepo.CountAssignmentInCourse(lesson.LessonID)
	countCompletedAssignment, _ := utilS.UserAssignmentRepo.CountAssignmentCompletedInCourse(lesson.LessonID, user.UserID)

	if countAssignment == countCompletedAssignment {
		userLesson, _ := utilS.UserLessonRepo.FindByUserIDAndLessonID(lesson.LessonID, user.UserID)
		userLesson.Status = true
		userLesson.CompletedAt = time.Now()
		if err := utilS.UserLessonRepo.Update(*userLesson); err != nil {
			repo.OutPutDebugError(err.Error())
		}
	}

	countLesson, _ := utilS.LessonRepo.CountLessonInCourse(lesson.CourseID)
	countCompletedLesson, _ := utilS.UserLessonRepo.CountLessonCompletedInCourse(lesson.CourseID, user.UserID)

	courseUser, _ := utilS.CourseUserRepo.FindByUserIDAndCourseID(lesson.CourseID, user.UserID)
	if countLesson != 0 && countLesson == countCompletedLesson {
		courseUser.Status = true
		courseUser.CompletedAt = time.Now()
		if err := utilS.CourseUserRepo.Update(*courseUser); err != nil {
			repo.OutPutDebugError(err.Error())
		}
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Student Assignment Submitted", fiber.Map{
		"FileAssignments": fileAssignments,
		"UserAssignment":  userAssignment,
		"CourseUser":      courseUser,
	})
}

func ApiReturnSubmitAssignment(c *fiber.Ctx) error {
	userAssignmentID := utilS.StringToUint32(c.Params("id"))
	if userAssignment, err := utilS.UserAssignmentRepo.FindByUserAssignmentID(userAssignmentID); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Can not undo assignment ", nil)
	} else {
		userAssignment.Status = false
		if err := utilS.UserAssignmentRepo.Update(*userAssignment); err != nil {
			repo.OutPutDebugError(err.Error())
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Can not undo assignment", nil)
		}
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Successfully", nil)
}
func ApiAssignTopicToAssignment(c *fiber.Ctx) error {
	type Request struct {
		AssignmentID uint32 `json:"assignment_id"`
		TopicID      uint32 `json:"topic_id"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid format", nil)
	}

	if _, err := utilS.TopicAssignmentRepo.FindTopicsByAssignmentIDTopicID(request.AssignmentID, request.TopicID); err == nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Only one topic can be selected for the assignment.", nil)
	}

	assignment, _ := utilS.AssignmentRepo.FindByID(request.AssignmentID)

	topic, err := utilS.TopicRepo.FindByID(request.TopicID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Topic not found", nil)
	}

	topicAssignment := models.TopicAssignment{
		AssignmentID: request.AssignmentID,
		TopicID:      request.TopicID,
		CreatedBy:    GetInfoUser(c).UserID,
		UpdatedBy:    GetInfoUser(c).UserID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	assignment.Score = topic.TotalScore
	if err := utilS.AssignmentRepo.Update(*assignment); err != nil {
		repo.OutPutDebugError(err.Error())
	}

	if err := utilS.TopicAssignmentRepo.Create(&topicAssignment); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "create successfully", topicAssignment)
}

func ApiGetTopicByAssignmentID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetTopicByAssignmentID")
	assignmentID := utilS.StringToUint32(c.Params("id"))
	topic, err := utilS.TopicAssignmentRepo.FindTopicsByAssignmentID(assignmentID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Topics retrieved successfully", topic)
}
func GetTopicQuestion(c *fiber.Ctx) error {
	repo.OutPutDebug("CreateTopic")
	return c.Render("pages/assignment/topic_question", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}
