package main

import (
	"github.com/gin-gonic/gin"

	"api/src/association/application"
	"api/src/association/infraestructure/adapters"
	"api/src/association/infraestructure/http/controllers"
	"api/src/association/infraestructure/http/routes"
	"api/src/membership/application"
	"api/src/membership/infraestructure/adapters"
	"api/src/membership/infraestructure/http/controllers"
	"api/src/membership/infraestructure/http/routes"
	"api/src/user/application"
	"api/src/user/infraestructure/adapters"
	"api/src/user/infraestructure/http/controllers"
	"api/src/user/infraestructure/http/routes"
)

func main() {
	myGin := gin.Default()

	userRepository := adapters.NewMySQLUserRepository()
	createUserUseCase := application.NewCreateUserUseCase(userRepository)
	createUserController := controllers.NewUserController(createUserUseCase)
	routes.RegisterUserRoutes(myGin, createUserController)

	membershipRepository := adapters.NewMySQLMembershipRepository()
	createMembershipUseCase := application.NewCreateMembershipUseCase(membershipRepository)
	createMembershipController := controllers.NewMembershipController(createMembershipUseCase)
	routes.RegisterMembershipRoutes(myGin, createMembershipController)

	associationRepository := adapters.NewMySQLAssociationRepository()
	createAssociationUseCase := application.NewCreateAssociationUseCase(associationRepository)
	createAssociationController := controllers.NewAssociationController(createAssociationUseCase)
	routes.RegisterAssociationRoutes(myGin, createAssociationController)

	myGin.Run()
}
