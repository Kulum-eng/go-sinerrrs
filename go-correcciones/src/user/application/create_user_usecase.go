package application

import (
	"api/src/user/domain"
	"api/src/user/domain/ports"
	"api/src/user/infraestructure/email" // ✅ importamos el simulador de correo
)

type CreateUserUseCase struct {
	repo ports.UserRepository
}

func NewCreateUserUseCase(repo ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

func (uc *CreateUserUseCase) Execute(user domain.User) (int, error) {
	id, err := uc.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	// ✅ Simulación de envío de correo
	email.SendWelcomeEmail(user.Email, user.Name)

	return id, nil
}

