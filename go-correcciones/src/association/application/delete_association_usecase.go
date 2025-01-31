package application

import (
	"api/src/association/domain/ports"
)

type DeleteAssociationUseCase struct {
	repo ports.AssociationRepository
}

func NewDeleteAssociationUseCase(repo ports.AssociationRepository) *DeleteAssociationUseCase {
	return &DeleteAssociationUseCase{repo: repo}
}

func (useCase *DeleteAssociationUseCase) Execute(id int) error {
	return useCase.repo.Delete(id)
}
