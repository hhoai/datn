package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"strings"
	"time"
)

func GetSkill(c *fiber.Ctx) error {
	repo.OutPutDebug("GetSkill")
	return c.Render("pages/skills/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}

func ApiGetSkill(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetSkill")
	skills := utilS.SkillRepo.FindAll()

	return utilS.ResultResponse(c, fiber.StatusOK, "Skills retrieved successfully", skills)

}
func ApiGetSkillByID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetSkillByID")
	skillID := utilS.StringToUint32(c.Params("id"))

	skill, err := utilS.SkillRepo.FindByID(skillID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Skill not found", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Skills retrieved successfully", skill)
}
func ApiCreateSkill(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateSkill")

	type Request struct {
		Name string `json:"name"`
	}
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid format", nil)
	}
	if strings.TrimSpace(request.Name) == "" {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Name field is required", nil)
	}
	user := GetInfoUser(c)
	skill := models.Skill{
		Name:      request.Name,
		CreatedBy: user.UserID,
	}

	if err := utilS.SkillRepo.Create(&skill); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusCreated, "Skills create successfully", skill)

}

func ApiDeleteSkill(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteSkill")

	skillID := utilS.StringToUint32(c.Params("id"))

	skill, err := utilS.SkillRepo.FindByID(skillID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Skill not found", nil)
	}
	if err = utilS.SkillRepo.Delete(skill); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to delete", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Skills deleted successful", nil)

}
func ApiUpdateSkill(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateSkill")
	skillID := utilS.StringToUint32(c.Params("id"))

	skill, err := utilS.SkillRepo.FindByID(skillID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "Skill not found", nil)
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

	skill.Name = request.Name
	skill.UpdatedAt = time.Now()
	user := GetInfoUser(c)
	skill.UpdatedBy = user.UserID

	if err := utilS.SkillRepo.Update(*skill); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to update", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Skills update", skill)
}
