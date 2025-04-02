package handlers

import (
	"github.com/gofiber/fiber/v2"
	"lms/models"
	"lms/repo"
	"lms/utilS"
	"time"
)

func GetRoles(c *fiber.Ctx) error {
	repo.OutPutDebug("GetUsers")
	return c.Render("pages/roles/index", fiber.Map{
		"Ctx": c,
	}, "layouts/main")

}

func ApiGetRoles(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetRoles")

	roles := utilS.RoleRepo.FindAll()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "roles retrieved successfully",
		"data":    roles,
	})
}

func ApiGetPermissions(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetPermissions")

	permissions := utilS.PermissionRepo.FindAll()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "permissions retrieved successfully",
		"data":    permissions,
	})

}
func ApiGetRoleByID(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiGetRoleByID")

	roleID := utilS.StringToUint32(c.Params("id"))

	role, err := utilS.RoleRepo.FindByID(roleID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    404,
			"message": "Role not found",
		})
	}

	permissions, err := utilS.RolePerRepo.FindByRoleID(roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Failed to retrieve permissions for the role",
		})
	}
	if len(permissions) == 0 {
		permissions = []models.RolePermission{}
	}

	allPermissions := utilS.PermissionRepo.FindAll()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "Role retrieved successfully",
		"data": fiber.Map{
			"role":           role,
			"permissions":    permissions,
			"allPermissions": allPermissions,
		},
	})
}

func ApiCreateRole(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiCreateRole")

	type Request struct {
		RoleName    string   `json:"name"`
		Permissions []uint32 `json:"permissions"`
	}
	var (
		request         = Request{}
		rolePermissions []*models.RolePermission
	)
	if err := c.BodyParser(&request); err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid request", nil)
	}
	user := GetInfoUser(c)
	role := models.Role{
		Name:      request.RoleName,
		CreatedBy: user.UserID,
		UpdatedBy: user.UserID,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	if err := utilS.RoleRepo.Create(&role); err != nil {
		if err := utilS.RolePerRepo.CreateAny(rolePermissions); err != nil {
			return utilS.ResultResponse(c, fiber.StatusAccepted, "Failed to create role", nil)
		}
	}

	for _, permissionID := range request.Permissions {
		rolePermissions = append(rolePermissions, &models.RolePermission{
			RoleID:       role.RoleID,
			PermissionID: permissionID,
			CreatedBy:    user.UserID,
			UpdatedBy:    user.UserID,
			UpdatedAt:    time.Now(),
			CreatedAt:    time.Now(),
		})
	}
	if err := utilS.RolePerRepo.CreateAny(rolePermissions); err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Failed to assign permission to role", nil)
	}

	return utilS.ResultResponse(c, fiber.StatusOK, "Role created successfully", nil)
}

func ApiDeleteRole(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiDeleteRole")

	roleID := utilS.StringToUint32(c.Params("id"))
	role, err := utilS.RoleRepo.FindByID(roleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Role not found",
		})
	}

	if err := utilS.RoleRepo.Delete(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Failed to delete role",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Role deleted successfully",
	})

}

func ApiUpdateRole(c *fiber.Ctx) error {
	repo.OutPutDebug("ApiUpdateRole")

	type Request struct {
		RoleName    string   `json:"name"`
		Permissions []uint32 `json:"permissions"`
	}
	var (
		request         Request
		rolePermissions []*models.RolePermission
	)
	if err := c.BodyParser(&request); err != nil {
		return utilS.ResultResponse(c, fiber.StatusBadRequest, "Invalid request", nil)
	}

	roleID := utilS.StringToUint32(c.Params("id"))
	role, err := utilS.RoleRepo.FindByID(roleID)
	if err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Role not found", nil)
	}
	role.Name = request.RoleName
	role.UpdatedAt = time.Now()
	user := GetInfoUser(c)
	role.UpdatedBy = user.UserID

	if err := utilS.RoleRepo.Update(*role); err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Failed to update role", nil)
	}

	if err := utilS.RolePerRepo.DeleteByRoleID(role.RoleID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "Failed to delete old permissions",
		})
	}

	for _, permissionID := range request.Permissions {
		rolePermissions = append(rolePermissions, &models.RolePermission{
			RoleID:       role.RoleID,
			PermissionID: permissionID,
			CreatedBy:    user.UserID,
			UpdatedBy:    user.UserID,
			UpdatedAt:    time.Now(),
			CreatedAt:    time.Now(),
		})

	}
	if err := utilS.RolePerRepo.CreateAny(rolePermissions); err != nil {
		return utilS.ResultResponse(c, fiber.StatusAccepted, "Failed to assign permission to role", nil)
	}
	return utilS.ResultResponse(c, fiber.StatusOK, "Role updated successfully", nil)
}
