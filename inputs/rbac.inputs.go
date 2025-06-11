package inputs


// CreateStaffRequest ...
type CreateStaffRequest struct {
	GarageID  string `json:"garageID" binding:"required"`
	UserEmail string `json:"userEmail" binding:"required"`
	RoleID    string `json:"roleID" binding:"required"`
}

// UpdateStaffRequest ...
type UpdateStaffRequest struct {
	RoleID string `json:"roleID" binding:"required"`
}

// AddRoletoStaffInput ...
type AddRoleToStaff struct {
	StaffID string `json:"staffID" binding:"required"`
	RoleID  string `json:"roleID" binding:"required"`
}