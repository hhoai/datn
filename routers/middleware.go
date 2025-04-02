package routers

import (
	"github.com/gofiber/fiber/v2"
	"lms/handles"
	"lms/utilS"
)

func ErrorMiddleware(ctx *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			_ = utilS.ErrorHandler(ctx, fiber.ErrInternalServerError)
		}
	}()
	return ctx.Next()
}

func PageErrorMiddleware(ctx *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			_ = utilS.PageErrorHandler(ctx, fiber.ErrInternalServerError)
		}
	}()
	return ctx.Next()
}

func IsAuthenticated(c *fiber.Ctx) error {
	if !handlers.CheckAuthentication(c) {
		return c.Redirect("/")
	}
	return c.Next()
}

func PermissionManagerUsers(c *fiber.Ctx) error {
	if !handlers.CheckPermission("manager_users", c) {
		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
	}
	return c.Next()
}
func PermissionManagerRoles(c *fiber.Ctx) error {
	if !handlers.CheckPermission("manager_roles", c) {
		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
	}
	return c.Next()
}

func PermissionManagerCourses(c *fiber.Ctx) error {
	if !handlers.CheckPermission("manager_courses", c) {
		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
	}
	return c.Next()
}

func PermissionManagerLessons(c *fiber.Ctx) error {
	if !handlers.CheckPermission("manager_lessons", c) {
		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
	}
	return c.Next()
}

func PermissionManagerQuestions(c *fiber.Ctx) error {
	if !handlers.CheckPermission("manager_questions", c) {
		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
	}
	return c.Next()
}

func PermissionManagerPrograms(c *fiber.Ctx) error {
	if !handlers.CheckPermission("manager_programs", c) {
		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
	}

	return c.Next()
}
func PermissionManagerTopics(c *fiber.Ctx) error {
	if !handlers.CheckPermission("manager_topics", c) {
		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
	}
	return c.Next()
}

//
//func CheckCourseStatus(c *fiber.Ctx) error {
//	courseID := c.Params("courseID")
//	course, err := utilS.CourseRepo.FindByID(utilS.StringToUint32(courseID))
//	if err != nil {
//		return utilS.PageErrorResponse(c, fiber.StatusBadRequest)
//	}
//
//	if !course.Status {
//		return utilS.PageErrorResponse(c, fiber.StatusForbidden)
//	}
//
//	return c.Next()
//}
