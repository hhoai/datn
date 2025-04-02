package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/repo"
	"lms/utilS"
)

func GetTypeUser(c *fiber.Ctx) error {
	repo.OutPutDebug(" GetTypeUser")
	return c.Render("pages/type_user/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")
}
func ApiGetTypeUsers(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetTypeUsers")

	typeUsers := utilS.TypeUserRepo.FindAll()

	return utilS.ResultResponse(c, fiber.StatusOK, "type user retrieved successfully", typeUsers)
}

func ApiGetTypeUserByID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetTypeUserByID")
	typeUserID := utilS.StringToUint32(c.Params("id"))

	typeUser, err := utilS.TypeUserRepo.FindByID(typeUserID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "type user not found", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "type user retrieved successfully", typeUser)
}
func ApiCreateTypeUser(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateTypeUser")

	//type Request struct {
	//	Name string `json:"name"`
	//}
	//var request Request
	//if err := c.BodyParser(&request); err != nil {
	//	return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid format", nil)
	//}
	//user := GetInfoUser(c)
	//typeUser := models.TypeUser{
	//	Name:      request.Name,
	//	CreatedBy: user.UserID,
	//}
	//
	//if err := utilS.TypeUserRepo.Create(&typeUser); err != nil {
	//	return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to create", nil)
	//}
	return utilS.ResultResponse(c, fiber.StatusCreated, "Type User create successfully", nil)

}

func ApiDeleteTypeUser(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteTypeUser")

	typeUserID := utilS.StringToUint32(c.Params("id"))

	typeUser, err := utilS.TypeUserRepo.FindByID(typeUserID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "type user not found", nil)
	}
	if err := utilS.TypeUserRepo.Delete(typeUser); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to delete", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Type User deleted successful", nil)

}
func ApiUpdateTypeUser(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateTypeUser")
	typeUserID := utilS.StringToUint32(c.Params("id"))

	typeUser, err := utilS.TypeUserRepo.FindByID(typeUserID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusNotFound, "type user not found", nil)
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

	typeUser.Name = request.Name
	//typeUser.UpdatedAt = time.Now()
	//user := GetInfoUser(c)
	//typeUser.UpdatedBy = user.UserID

	if err := utilS.TypeUserRepo.Update(*typeUser); err != nil {
		return utilS.ResultResponse(c, fiber.StatusInternalServerError, "Failed to update", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Type user updated ok", typeUser)
}
