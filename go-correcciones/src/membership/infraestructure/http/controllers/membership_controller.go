package controllers

import (
	"api/src/membership/application"
	"api/src/membership/domain"
	"api/src/user/infraestructure/http/responses"

	"net/http"

	"github.com/gin-gonic/gin"
)

type MembershipController struct {
	useCase *application.CreateMembershipUseCase
}

func NewMembershipController(useCase *application.CreateMembershipUseCase) *MembershipController {
	return &MembershipController{useCase: useCase}
}

func (ctrl *MembershipController) CreateMembership(ctx *gin.Context) {
	var membership domain.Membership
	if err := ctx.ShouldBindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	if err := ctrl.useCase.Execute(membership); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear la membresía", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Membresía creada exitosamente", membership))
}
