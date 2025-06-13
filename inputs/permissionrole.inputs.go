package inputs

// AddPermissionToRole ..
type AddPermissionToRoleInput struct {
	RoleID string `json:"roleID"`
	//Permissions comma  separated list of applicable permissions
	Permissions string `json:"permissions"`
}
