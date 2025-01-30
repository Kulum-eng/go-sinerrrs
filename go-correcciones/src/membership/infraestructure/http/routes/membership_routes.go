package routes

import (
	"api/src/membership/infraestructure/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMembershipRoutes(router *gin.Engine, controller *controllers.MembershipController) {
	membershipRoutes := router.Group("/membership")
	{
		membershipRoutes.POST("/", controller.CreateMembership)
		/* membershipRoutes.GET("/:id", controller.GetMembership)
		membershipRoutes.GET("/", controller.GetAllMemberships)
		membershipRoutes.PUT("/:id", controller.UpdateMembership)
		membershipRoutes.DELETE("/:id", controller.DeleteMembership) */
	}
}
