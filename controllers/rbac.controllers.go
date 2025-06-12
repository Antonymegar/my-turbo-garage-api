package controllers

import (
	"fmt"
	"myturbogarage/config"
	"myturbogarage/helpers"
	"myturbogarage/inputs"
	"myturbogarage/models"
	"myturbogarage/outputs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// getUserFromContext ...
func getUserFromContext(ctx *gin.Context) (*helpers.AuthClaims, error) {
	val, ok := ctx.Get("claims")
	if !ok {
		return nil, fmt.Errorf("could not get user from context")
	}

	user, ok := val.(*helpers.AuthClaims)
	if !ok {
		return nil, fmt.Errorf("could not parse user from context")
	}

	return user, nil
}

// WhoAmI loads the user's profile based on th platform the user is on
func WhoAmI(ctx *gin.Context) error {
	ctxUser, err := getUserFromContext(ctx)
	if err != nil {
		return err
	}

	user, err := config.FindOne(&models.User{ID: ctxUser.ID})
	if err != nil {
		return err
	}

	lg := outputs.NewLoginResponse(user)

	ctx.JSON(200, lg)
	return nil
}

// CreateStaff ...
func CreateStaff(c *gin.Context) error {
	var req inputs.CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}

	var userID string
	result := config.Conn().Raw("SELECT id FROM users WHERE email = ?", req.UserEmail).Scan(&userID)

	if result.Error != nil {
		// TODO if no user create the user
		return result.Error
	}

	staff := models.Staff{
		ID:       uuid.NewString(),
		RoleID:   req.RoleID,
		GarageID: req.GarageID,
		UserID:   userID,
	}

	// TODO send invite email to this user
	if err := config.Create(&staff); err != nil {
		return err
	}

	c.JSON(http.StatusOK, staff)
	return nil
}

// UpdateStaff ...
func UpdateStaff(c *gin.Context) error {
	var req inputs.UpdateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}

	staff, err := config.FindOne(&models.Staff{ID: c.Param("id")})
	if err != nil {
		return err
	}

	if err := config.Update(&staff); err != nil {
		return err
	}

	return nil
}

// DeleteStaff ...
func DeleteStaff(c *gin.Context) error {
	staff, err := config.FindOne(&models.Staff{ID: c.Param("id")})
	if err != nil {
		return err
	}

	if err := config.Delete(&staff); err != nil {
		return err
	}

	return nil
}

// GetStaffs ...
func GetStaffs(c *gin.Context) error {
	garageID := c.Query("garageID")
	// Validate that the vendorID is provided
	if garageID == "" {
		return fmt.Errorf("garageID is required")
	}
	// Find all staff members associated with the specified vendorID
	staffList, err := config.FindAll(&models.Staff{GarageID: garageID})
	if err != nil {
		return err
	}

	staffDetails := []*struct {
		Staff *models.Staff
		Role  *models.Role
		User  *models.User
	}{}

	for _, staff := range staffList {
		// Find the role associated with the staff member
		role, err := config.FindOne(&models.Role{ID: staff.RoleID})
		if err != nil {
			return err
		}

		// Find the user associated with the staff member
		user, err := config.FindOne(&models.User{ID: staff.UserID})
		if err != nil {
			return err
		}

		staffDetails = append(staffDetails, &struct {
			Staff *models.Staff
			Role  *models.Role
			User  *models.User
		}{
			Staff: staff,
			Role:  role,
			User:  user,
		})
	}

	c.JSON(http.StatusOK, staffDetails)
	return nil
}

// CreateRole ...
func CreateRole(c *gin.Context) error {
	var req inputs.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}
	/// check role of similar name exists
	count, err := config.Count(&models.Role{Name: req.Name})
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("role %s  in this vendor already exists", req.Name)
	}

	role := models.Role{
		ID:          uuid.NewString(),
		GarageID:    req.GarageID,
		Name:        req.Name,
		Description: req.Description,
		Permissions: req.Permissions,
	}
	if err := config.Create(&role); err != nil {
		return err
	}

	return nil
}

// UpdateRole ...
func UpdateRole(c *gin.Context) error {
	var req inputs.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}

	role, err := config.FindOne(&models.Role{ID: c.Param("id")})
	if err != nil {
		return err
	}

	role.GarageID = req.GarageID
	role.Name = req.Name
	role.Description = req.Description
	role.Permissions = req.Permissions

	if err := config.Update(&role); err != nil {
		return err
	}

	return nil
}

// DeleteRole ...
func DeleteRole(c *gin.Context) error {
	role, err := config.FindOne(&models.Role{ID: c.Param("id")})
	if err != nil {
		return err
	}

	if err := config.Delete(&role); err != nil {
		return err
	}

	return nil
}

// GetRoles ...
func GetRoles(c *gin.Context) error {
	garageID := c.Query("garageID")
	roles, err := config.FindAll(&models.Role{GarageID: garageID})
	if err != nil {
		return err
	}

	c.JSON(200, roles)
	return nil
}

// AddRoleToStaff
func AddRoleToStaff(ctx *gin.Context) error {
	var req inputs.AddRoleToStaff

	staff, err := config.FindOne(&models.Staff{ID: req.StaffID})
	if err != nil {
		return err
	}
	role, err := config.FindOne(&models.Role{ID: req.RoleID})
	if err != nil {
		return err
	}
	staff.RoleID = role.ID
	if err := config.Update(&staff); err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Role assigned to staff successfully"})
	return nil
}
