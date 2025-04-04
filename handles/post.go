package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/structs"
	"lms/utilS"
	"log"
	"time"
)

func ApiGetPostByLessonID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetPosts")
	lessonID := utilS.StringToUint32(c.Params("id"))

	posts, _ := utilS.PostRepo.FindByLessonID(lessonID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Posts retrieved successfully",
		"data":    posts,
	})

}

func ApiGetPosts(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetCourses")

	posts := utilS.PostRepo.FindAll()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Posts retrieved successfully",
		"data":    posts,
	})
}

func GetStudentPostDetails(c *fiber.Ctx) error {
	repo.OutPutDebug("GetStudentPostDetails")
	lessonID := utilS.StringToUint32(c.Params("lesson_id"))
	p := c.Params("id")

	return c.Render("pages/post/student-post", fiber.Map{
		"Ctx":      c,
		"PostID":   p,
		"LessonID": lessonID,
	}, "layouts/main")
}

func GetPostID(c *fiber.Ctx) error {
	repo.OutPutDebug("GetPostID")
	p := c.Params("id")
	return c.Render("pages/post/detail", fiber.Map{
		"Ctx":    c,
		"PostID": p,
	}, "layouts/main")
}

func ApiGetStudentPostDetails(c *fiber.Ctx) error {
	repo.OutPutDebug("GetStudentPostDetails")
	p := c.Params("id")

	postID := utilS.StringToUint32(p)

	post, err := utilS.PostRepo.FindByID(postID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
		return err
	}

	lesson, _ := utilS.LessonRepo.FindByID(post.LessonID)

	course, _ := utilS.CourseRepo.FindByID(lesson.CourseID)

	courseStatus := course.Status

	file, _ := utilS.FilePostRepo.FindByPostID(postID)

	//comments, _ := utilS.CommentRepo.FindByPostID(postID)

	user := GetInfoUser(c)

	return utilS.ResultResponse(c, fiber.StatusOK, "get post success", fiber.Map{
		"CreatedBy":    post.User.Name,
		"CreatedAt":    post.CreatedAt,
		"PostTitle":    post.Title,
		"PostBody":     post.Body,
		"Files":        file,
		"CourseTitle":  course.Title,
		"CourseStatus": courseStatus,
		"LessonID":     post.LessonID,
		"Title":        lesson.Title,
		//"Comments":     comments,
		"Username": user.Name,
	})
}

func CreatePost(c *fiber.Ctx) error {
	repo.OutPutDebug("CreatePost")

	var form structs.CreatePostForm
	if err := c.BodyParser(&form); err != nil {
		repo.OutPutDebugError("Failed to parse body: " + err.Error())
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Incorrect format", nil)
	}

	var newPost models.Post
	var user models.UserWithoutPass
	user = GetInfoUser(c)

	newPost.Title = form.Title
	newPost.Body = form.Body
	newPost.LessonID = utilS.StringToUint32(form.LessonID)
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()
	newPost.CreatedBy = user.UserID
	newPost.UpdatedBy = user.UserID

	if err := utilS.PostRepo.Create(newPost); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Error creating post", nil)
	}

	lesson, _ := utilS.LessonRepo.FindByID(newPost.LessonID)

	courseUsers := utilS.CourseUserRepo.FindUserByCourseID(lesson.CourseID, 4)

	var studentEmail []string
	for _, student := range courseUsers {
		if student.User.TypeUserID == 4 {
			studentEmail = append(studentEmail, student.User.Email)
		}
	}

	title := "Thông báo lớp học"
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
			<p>Giáo viên vừa đăng tải một bài viết mới trên hệ thống.</p>
		</div>
	</body>
	</html>
`

	err := SendEmailToStudents(studentEmail, title, body)
	if err != nil {
		log.Println("Failed to send email:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email"})
	}

	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Post created successfully",
		Data:    nil,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	repo.OutPutDebug("UpdatePost")
	var user models.UserWithoutPass
	user = GetInfoUser(c)

	postID := utilS.StringToUint32(c.Params("id"))

	post, _ := utilS.PostRepo.FindByID(postID)

	var form structs.UpdatePostForm

	form.Title = c.FormValue("title")
	form.Body = c.FormValue("body")
	formData, err := c.MultipartForm()
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Incorrect form data", nil)
	}

	fileIds := formData.Value["fileIds[]"]

	if err := utilS.FilePostRepo.DeleteDefault(); err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Can not update file post", nil)
	}

	if len(fileIds) != 0 {
		if err := utilS.FilePostRepo.UpdateDefaultByFileIDs(fileIds, true); err != nil {
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Can not update file post", nil)
		}
	}

	filepath := "./documents/posts/" + utilS.Uint32ToString(post.LessonID)
	if err := utilS.CreateFolder(filepath); err != nil {
		repo.OutPutDebugError(err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		repo.OutPutDebugError(err.Error())

	} else {
		err = c.SaveFile(file, filepath+"/"+file.Filename)
		if err != nil {
			//return utilS.ResultResponse(c, fiber.StatusNotAcceptable, "Can not update file", nil)
			repo.OutPutDebugError(err.Error())
		}

		var filePost models.FilePost
		filePost.Filename = file.Filename
		filePost.PostID = postID

		if err := utilS.FilePostRepo.Create(filePost); err != nil {
			repo.OutPutDebugError(err.Error())
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Can not update file", nil)
		}
	}

	post.Title = form.Title
	post.Body = form.Body
	post.UpdatedBy = user.UserID
	post.UpdatedAt = time.Now()

	if err := utilS.PostRepo.Update(*post); err != nil {
		repo.OutPutDebugError(err.Error())
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Can not update assignment", nil)
	}

	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Update successfully",
		Data:    nil,
	})
}

func GetPostByID(c *fiber.Ctx) error {
	postID := utilS.StringToUint32(c.Params("id"))

	post, err := utilS.PostRepo.FindByID(postID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Get post successfully",
		Data:    fiber.Map{"post": post},
	})
}

func ApiGetPostComment(c *fiber.Ctx) error {

	repo.OutPutDebug("ApiGetPostComment")
	p := c.Params("id")

	postID := utilS.StringToUint32(p)

	post, err := utilS.PostRepo.FindByID(postID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
		return err
	}

	lesson, _ := utilS.LessonRepo.FindByID(post.LessonID)

	file, _ := utilS.FilePostRepo.FindByPostID(postID)

	//comments, _ := utilS.CommentRepo.FindByPostID(postID)

	user := GetInfoUser(c)

	return utilS.ResultResponse(c, fiber.StatusOK, "get post success", fiber.Map{
		"CreatedBy": post.User.Name,
		"CreatedAt": post.CreatedAt,
		"PostTitle": post.Title,
		"PostBody":  post.Body,
		"Files":     file,
		"LessonID":  post.LessonID,
		"Title":     lesson.Title,
		//"Comments":  comments,
		"Username": user.Name,
	})

}

func DeletePost(c *fiber.Ctx) error {
	repo.OutPutDebug("DeletePost")
	postID := utilS.StringToUint32(c.Params("id"))

	err := utilS.PostRepo.DeleteByID(postID)
	if err != nil {
		repo.OutPutDebugError(err.Error())
	}
	return c.JSON(&utilS.Response{
		Code:    fiber.StatusOK,
		Message: "Post deleted successfully",
		Data:    nil,
	})
}
