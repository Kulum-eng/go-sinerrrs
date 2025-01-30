package application

import (
	"api/src/user/domain"
	"api/src/user/domain/ports"
)

type UpdateUserUseCase struct {
	repo ports.UserRepository
}

func NewUpdateUserUseCase(repo ports.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo}
}

func (uc *UpdateUserUseCase) Execute(user domain.User) error {
	return uc.repo.UpdateUser(user)
}
