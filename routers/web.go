package routers

import (
	"github.com/gofiber/fiber/v2"
	handlers "lms/handles"
)

func Router(app *fiber.App) {
	app.Use(PageErrorMiddleware)
	app.Get("", handlers.GetLogin)
	app.Get("/login", handlers.GetLogin)
	app.Get("/signup", handlers.GetSignup)
	app.Post("/logout", handlers.GetLogout)
	app.Get("/logout", handlers.GetLogout)
	app.Get("/reset-password/:id/:token", handlers.ApiPostActivationPassword)
	app.Get("/update-email/:id/:token/:email", handlers.ApiConfirmEmail)
	app.Get("/auth/google", handlers.GoogleLogin)
	app.Get("/auth/google/callback", handlers.GoogleCallback)

	app.Get("/dashboard", IsAuthenticated, handlers.GetDashboard)

	users := app.Group("/users", IsAuthenticated, PermissionManagerUsers)
	users.Get("", handlers.GetUsers)
	users.Post("/create", handlers.CreateUser)
	users.Put("/update/:id", handlers.UpdateUser)
	users.Get("/:id/courses", handlers.GetDetailUser)

	courses := app.Group("/courses", IsAuthenticated, PermissionManagerCourses)
	courses.Get("", handlers.GetCourses)
	courses.Get("/:id", handlers.GetCourseDetails)
	courses.Get("/:user_id/assignments/:assignment_id/topics/:id", handlers.GetStudentAssignmentCourseTopicID)

	courseCategories := app.Group("/course_categories", IsAuthenticated, PermissionManagerCourses)
	courseCategories.Get("", handlers.GetCourseCategory)

	lessons := app.Group("/lessons", IsAuthenticated, PermissionManagerLessons)
	lessons.Get("/:id/assignments", handlers.GetLessonDetails)
	lessons.Get("/:id/posts", handlers.GetPosts)
	//lessons.Get("/assignment/:assignment_id/scoring", handlers.GetScoringAssignment)

	levels := app.Group("/levels", IsAuthenticated, PermissionManagerQuestions)
	levels.Get("", handlers.GetLevel)

	assignments := app.Group("/assignments", IsAuthenticated, PermissionManagerLessons)
	assignments.Get("/:id/courses/:course_id", handlers.GetAssignmentDetails)
	assignments.Get("/:id/topic_question", handlers.GetTopicQuestion)

	lessonCategory := app.Group("/lesson_categories", IsAuthenticated, PermissionManagerLessons)
	lessonCategory.Get("", handlers.GetLessonCategory)

	typeUsers := app.Group("type_users", IsAuthenticated, PermissionManagerUsers)
	typeUsers.Get("", handlers.GetTypeUser)

	skills := app.Group("/skills", IsAuthenticated, PermissionManagerQuestions)
	skills.Get("", handlers.GetSkill)

	posts := app.Group("/posts", IsAuthenticated, PermissionManagerLessons)
	posts.Get("/:id", handlers.GetPostID)

	programs := app.Group("/programs", IsAuthenticated, PermissionManagerPrograms)
	programs.Get("", handlers.GetPrograms)
	programs.Get("/:id/details", handlers.GetProgramDetails)

	roles := app.Group("/roles", IsAuthenticated, PermissionManagerRoles)
	roles.Get("", handlers.GetRoles)

	information := app.Group("/information", IsAuthenticated)
	information.Get("", handlers.GetUserInformation)

	challenges := app.Group("/challenges", IsAuthenticated)
	challenges.Get("", handlers.GetChallenge)
}

func API(app *fiber.App) {
	apiV1 := app.Group("/api/v1", ErrorMiddleware)
	apiV1.Post("/login", handlers.ApiPostLogin)
	apiV1.Post("/signup", handlers.ApiPostSignup)
	apiV1.Post("/logout", handlers.ApiPostLogout)
	apiV1.Post("/forget-password", handlers.ApiPostForgetPassword)
	apiV1.Post("/resend-activation", handlers.ApiPostActivation)
	apiV1.Get("/resend-activation/:id/token/:token", handlers.CheckToken)

	users := apiV1.Group("/users", IsAuthenticated)
	users.Get("", handlers.ApiGetUsers)
	users.Get("/:id", handlers.ApiGetUserByID)
	users.Put("/:id", handlers.ApiUpdateUsers)
	users.Post("/create", handlers.ApiCreateUsers)
	users.Delete("/:id", handlers.ApiDeleteUser)
	users.Put("/:id/change_password", handlers.ApiChangePassword)
	users.Post("/fake-data-user", handlers.FakeDataUser)

	roles := apiV1.Group("/roles", IsAuthenticated)
	roles.Get("", handlers.ApiGetRoles)
	roles.Post("create", handlers.ApiCreateRole)
	roles.Delete("/:id", handlers.ApiDeleteRole)
	roles.Get("/:id", handlers.ApiGetRoleByID)
	roles.Put("/:id", handlers.ApiUpdateRole)

	typeUsers := apiV1.Group("/type_users", IsAuthenticated)
	typeUsers.Get("", handlers.ApiGetTypeUsers)
	typeUsers.Get("/:id", handlers.ApiGetTypeUserByID)
	typeUsers.Put("/:id", handlers.ApiUpdateTypeUser)
	typeUsers.Post("", handlers.ApiCreateTypeUser)
	typeUsers.Delete("/:id", handlers.ApiDeleteTypeUser)

	permissions := apiV1.Group("/permissions", IsAuthenticated)
	permissions.Get("", handlers.ApiGetPermissions)

	courses := apiV1.Group("/courses", IsAuthenticated)
	courses.Get("", handlers.ApiGetCourses)
	courses.Get("/:id", handlers.ApiGetCourseById)
	courses.Post("", handlers.ApiCreateCourse)
	courses.Put("/:id", handlers.ApiUpdateCourse)
	courses.Delete("", handlers.DeleteMultipleCourses)
	courses.Delete("/:id", handlers.ApiDeleteCourse)

	programs := apiV1.Group("/programs", IsAuthenticated)
	programs.Get("", handlers.ApiGetPrograms)
	programs.Get("/:id", handlers.ApiGetProgramByID)
	programs.Delete("/:id", handlers.ApiDeleteProgram)
	programs.Put("/:id", handlers.ApiUpdateProgram)
	programs.Post("", handlers.ApiCreateProgram)
	programs.Get("/prerequisite-course/:id", handlers.ApiGetPrerequisiteCourse)
	programs.Get("/:id/courses", handlers.ApiGetCourseInProgram)

	studentProgram := apiV1.Group("/student_programs", IsAuthenticated)
	studentProgram.Get("", handlers.ApiGetStudentProgram)

	requestProgram := apiV1.Group("/request-program", IsAuthenticated)
	requestProgram.Post("", handlers.PostReqJoinProgram)

	courseCategories := apiV1.Group("/course_categories", IsAuthenticated)
	courseCategories.Get("", handlers.ApiGetCourseCategory)
	courseCategories.Post("", handlers.ApiCreateCourseCategory)
	courseCategories.Delete("/:id", handlers.ApiDeleteCourseCategory)
	courseCategories.Put("/:id", handlers.ApiUpdateCourseCategory)
	courseCategories.Get("/:id", handlers.ApiGetCourseCategoryByID)

	levels := apiV1.Group("/levels", IsAuthenticated)
	levels.Get("", handlers.ApiGetLevel)
	levels.Get("/:id", handlers.ApiGetLevelByID)
	levels.Delete("/:id", handlers.ApiDeleteLevel)
	levels.Put("/:id", handlers.ApiUpdateLevel)
	levels.Post("", handlers.ApiCreateLevel)

	challenges := apiV1.Group("/challenges", IsAuthenticated)
	challenges.Get("", handlers.ApiGetChallenge)
	challenges.Put("/:id", handlers.ApiUpdateChallenge)
	challenges.Delete("/:id", handlers.ApiDeleteChallenge)
	challenges.Get("/:id", handlers.ApiGetChallengeByID)
	challenges.Post("", handlers.ApiCreateChallenge)

	lessons := apiV1.Group("/lessons", IsAuthenticated)
	lessons.Get("", handlers.ApiGetLessons)
	lessons.Get("/details/:id", handlers.GetLessonByID)
	lessons.Get("/:id", handlers.ApiGetLessonByCourseID)
	lessons.Post("", handlers.CreateLesson)
	lessons.Put("/:id", handlers.UpdateLesson)
	lessons.Delete("/:id", handlers.DeleteLesson)

	lessonCategories := apiV1.Group("/lesson-categories", IsAuthenticated)
	lessonCategories.Get("", handlers.ApiGetLessonCategory)
	lessonCategories.Post("", handlers.ApiCreateLessonCategory)
	lessonCategories.Put("/:id", handlers.ApiUpdateLessonCategory)
	lessonCategories.Delete("/:id", handlers.ApiDeleteLessonCategory)
	lessonCategories.Get("/:id", handlers.ApiGetLessonCategoryByID)

	assignments := apiV1.Group("/assignments", IsAuthenticated)
	assignments.Get("", handlers.ApiGetAssignments)
	assignments.Get("/:id", handlers.ApiGetAssignmentByLessonID)
	assignments.Get("/details/:id", handlers.GetAssignmentByID)
	assignments.Post("", handlers.CreateAssignment)
	assignments.Put("/upload/:id", handlers.UpdateAssignment)
	assignments.Delete("/:id", handlers.DeleteAssignment)
	assignments.Post("/assign_topic", handlers.ApiAssignTopicToAssignment)
	assignments.Get("/:id/topic", handlers.ApiGetTopicByAssignmentID)

	posts := apiV1.Group("/posts", IsAuthenticated)
	posts.Get("", handlers.ApiGetPosts)
	posts.Get("/:id", handlers.ApiGetPostByLessonID)
	posts.Get("/details/:id", handlers.GetPostByID)
	posts.Post("", handlers.CreatePost)
	posts.Put("/:id", handlers.UpdatePost)
	posts.Delete("/:id", handlers.DeletePost)
	posts.Get("/comment/:id", handlers.ApiGetPostComment)

	skills := apiV1.Group("/skills", IsAuthenticated)
	skills.Get("", handlers.ApiGetSkill)
	skills.Get("/:id", handlers.ApiGetSkillByID)
	skills.Delete("/:id", handlers.ApiDeleteSkill)
	skills.Put("/:id", handlers.ApiUpdateSkill)
	skills.Post("", handlers.ApiCreateSkill)

	information := apiV1.Group("/information", IsAuthenticated)
	information.Put("", handlers.ChangeEmail)
	information.Put("/change-password", handlers.PutUserInformation)

	dashboard := apiV1.Group("/dashboard", IsAuthenticated)
	dashboard.Get("", handlers.ApiGetDashboard)

	apiV1.Get("/search/:q", IsAuthenticated, handlers.ApiSearchPrograms)
}
