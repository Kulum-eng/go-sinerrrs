package application

import (
	"api/src/user/domain/ports"
)

type DeleteUserUseCase struct {
	repo ports.UserRepository
}

func NewDeleteUserUseCase(repo ports.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

func (uc *DeleteUserUseCase) Execute(id int) error {
	_, err := uc.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	return uc.repo.DeleteUser(id)
}
