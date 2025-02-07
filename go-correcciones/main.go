package main

import (
	"log"

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
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, GET, PUT")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
		c.Header("X-Frame-Options", "DENY")

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(CORS())

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
	u_routes.SetupUserRoutes(router, createUserController)

	membershipRepository := m_adapters.NewMySQLMembershipRepository(db)
	createMembershipUseCase := m_application.NewCreateMembershipUseCase(membershipRepository)
	getMembershipUseCase := m_application.NewGetMembershipUseCase(membershipRepository)
	updateMembershipUseCase := m_application.NewUpdateMembershipUseCase(membershipRepository)
	deleteMembershipUseCase := m_application.NewDeleteMembershipUseCase(membershipRepository)

	createMembershipController := m_controllers.NewMembershipController(createMembershipUseCase, getMembershipUseCase, updateMembershipUseCase, deleteMembershipUseCase)
	m_routes.RegisterMembershipRoutes(router, createMembershipController)

	associationRepository := a_adapters.NewMySQLAssociationRepository(db)
	createAssociationUseCase := a_application.NewCreateAssociationUseCase(associationRepository)
	getAssociationUseCase := a_application.NewGetAssociationUseCase(associationRepository)
	updateAssociationUseCase := a_application.NewUpdateAssociationUseCase(associationRepository)
	deleteAssociationUseCase := a_application.NewDeleteAssociationUseCase(associationRepository)

	createAssociationController := a_controllers.NewAssociationController(createAssociationUseCase, getAssociationUseCase, updateAssociationUseCase, deleteAssociationUseCase)
	a_routes.SetupRoutes(router, createAssociationController)

	router.Run()
}
