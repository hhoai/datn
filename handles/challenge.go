package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"time"
)

func GetChallenge(c *fiber.Ctx) error {
	repo.OutPutDebug("GetLevel")
	return c.Render("pages/challenges/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}

func ApiGetChallenge(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetChallenge")
	challenges := utilS.ChallengeRepo.FindAll()

	return utilS.ResultResponse(c, fiber.StatusOK, "Skills retrieved successfully", challenges)

}

func ApiGetChallengeByID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetChallengeByID")
	challengeID := utilS.StringToUint32(c.Params("id"))

	challenge, err := utilS.ChallengeRepo.FindByID(challengeID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "challenge not found", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "challenges retrieved successfully", challenge)
}
func ApiCreateChallenge(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateChallenge")

	type Request struct {
		Name string `json:"name"`
	}
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid format", nil)
	}
	user := GetInfoUser(c)
	challenge := models.Challenge{
		Name:      request.Name,
		CreatedBy: user.UserID,
	}

	if err := utilS.ChallengeRepo.Create(&challenge); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusCreated, "Skills create successfully", challenge)

}

func ApiDeleteChallenge(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteChallenge")

	challengeID := utilS.StringToUint32(c.Params("id"))

	challenge, err := utilS.ChallengeRepo.FindByID(challengeID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "challenge not found", nil)
	}
	if err := utilS.ChallengeRepo.Delete(challenge); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to delete", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Skills deleted successful", nil)

}
func ApiUpdateChallenge(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateSkill")
	challengeID := utilS.StringToUint32(c.Params("id"))

	challenge, err := utilS.ChallengeRepo.FindByID(challengeID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "challenge not found", nil)
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

	challenge.Name = request.Name
	challenge.UpdatedAt = time.Now()
	user := GetInfoUser(c)
	challenge.UpdatedBy = user.UserID

	if err := utilS.ChallengeRepo.Update(*challenge); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to update", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Challenge updated ok", challenge)
}
