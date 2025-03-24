package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"lms/repo"
	"strings"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
		StrictRouting:     true,
		CaseSensitive:     true,
		UnescapePath:      true,
		BodyLimit:         50 * 1024 * 1024, // 50MB
	})

	app.Use(func(c *fiber.Ctx) error {
		if strings.Contains(c.Path(), "..") ||
			strings.Contains(strings.ToLower(c.Path()), "script") ||
			strings.Contains(c.Path(), "(") ||
			strings.Contains(c.Path(), ";") ||
			strings.Contains(c.Path(), "--") {
			repo.OutPutDebugError(c.IP() + c.Path())
			return c.Redirect("/403")
		}
		return c.Next()
	})

	app.Static("/", "./public")
	app.Static("", "./documents")

	app.Listen(":8000")
}
