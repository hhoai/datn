package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"time"
)

func GetLevel(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLevel")
	return c.Render("pages/levels/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}

func ApiGetLevel(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLevel")

	levels := utilS.LevelRepo.FindAll()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Levels retrieved successfully",
		"data":    levels,
	})

}
func ApiGetLevelByID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetLevelByID")
	levelID := utilS.StringToUint32(c.Params("id"))

	level, err := utilS.LevelRepo.FindByID(levelID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "level not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"level":   level,
	})
}
func ApiCreateLevel(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateLevel")

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
	level := models.Level{
		Name:      request.Name,
		CreatedBy: user.UserID,
	}

	if err := utilS.LevelRepo.Create(&level); err != nil {
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

func ApiDeleteLevel(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteProgram")

	levelID := utilS.StringToUint32(c.Params("id"))

	level, err := utilS.LevelRepo.FindByID(levelID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "level not found",
		})
	}

	if err := utilS.LevelRepo.Delete(level); err != nil {
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

func ApiUpdateLevel(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateLevel")
	levelID := utilS.StringToUint32(c.Params("id"))

	level, err := utilS.LevelRepo.FindByID(levelID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "level not found",
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

	level.Name = request.Name
	level.UpdatedAt = time.Now()
	user := GetInfoUser(c)
	level.UpdatedBy = user.UserID

	if err := utilS.LevelRepo.Update(*level); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Failed to update level",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "level updated successfully",
	})
}
