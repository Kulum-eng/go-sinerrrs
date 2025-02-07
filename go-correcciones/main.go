package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	a_application "api/src/association/application"
	a_adapters "api/src/association/infraestructure/adapters"
	a_controllers "api/src/association/infraestructure/http/controllers"
	a_routes "api/src/association/infraestructure/http/routes"
	"api/src/core"

	m_application "api/src/membership/application"
	m_adapters "api/src/membership/infraestructure/adapters"
	m_controllers "api/src/membership/infraestructure/http/controllers"
	m_routes "api/src/membership/infraestructure/http/routes"
	u_application "api/src/user/application"
	u_adapters "api/src/user/infraestructure/adapters"
	u_controllers "api/src/user/infraestructure/http/controllers"
	u_routes "api/src/user/infraestructure/http/routes"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	myGin := gin.Default()

	myGin.Use(CORS())

	db, err := core.InitDB()
	if err != nil {
		log.Println(err)
		return
	}

	userRepository := u_adapters.NewMySQLUserRepository(db)
	createUserUseCase := u_application.NewCreateUserUseCase(userRepository)
	getUserUseCase := u_application.NewGetUserUseCase(userRepository)
	updateUserUseCase := u_application.NewUpdateUserUseCase(userRepository)
	deleteUserUseCase := u_application.NewDeleteUserUseCase(userRepository)

	createUserController := u_controllers.NewUserController(createUserUseCase, getUserUseCase, updateUserUseCase, deleteUserUseCase)
	u_routes.SetupUserRoutes(myGin, createUserController)

	membershipRepository := m_adapters.NewMySQLMembershipRepository(db)
	createMembershipUseCase := m_application.NewCreateMembershipUseCase(membershipRepository)
	getMembershipUseCase := m_application.NewGetMembershipUseCase(membershipRepository)
	updateMembershipUseCase := m_application.NewUpdateMembershipUseCase(membershipRepository)
	deleteMembershipUseCase := m_application.NewDeleteMembershipUseCase(membershipRepository)

	createMembershipController := m_controllers.NewMembershipController(createMembershipUseCase, getMembershipUseCase, updateMembershipUseCase, deleteMembershipUseCase)
	m_routes.RegisterMembershipRoutes(myGin, createMembershipController)

	associationRepository := a_adapters.NewMySQLAssociationRepository(db)
	createAssociationUseCase := a_application.NewCreateAssociationUseCase(associationRepository)
	getAssociationUseCase := a_application.NewGetAssociationUseCase(associationRepository)
	updateAssociationUseCase := a_application.NewUpdateAssociationUseCase(associationRepository)
	deleteAssociationUseCase := a_application.NewDeleteAssociationUseCase(associationRepository)

	createAssociationController := a_controllers.NewAssociationController(createAssociationUseCase, getAssociationUseCase, updateAssociationUseCase, deleteAssociationUseCase)
	a_routes.SetupRoutes(myGin, createAssociationController)

	myGin.Run()
}
