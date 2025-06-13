package routes

import (
	"myturbogarage/controllers"

	"github.com/gin-gonic/gin"
)

func initRBAC(router *gin.Engine) {
	r := router.Group("/rbac")
	{
		// roles
		r.POST("/roles", authCheck(controllers.CreateRole))
		r.PUT("/roles/:id", authCheck(controllers.UpdateRole))
		r.DELETE("/roles/:id", authCheck(controllers.DeleteRole))
		r.GET("/roles", authCheck(controllers.GetRoles))
		r.POST("/addRoleToStaff", authCheck(controllers.AddRoleToStaff))
		r.POST("/addPermissionToRole", authCheck(controllers.AddPermissionToRole))

		// staff
		r.POST("/staffs", authCheck(controllers.CreateStaff))
		r.GET("/staffs", authCheck(controllers.GetStaffs))
		r.PUT("/staffs/:id", authCheck(controllers.UpdateStaff))
		r.DELETE("/staffs/:id", authCheck(controllers.DeleteStaff))
	}
}
