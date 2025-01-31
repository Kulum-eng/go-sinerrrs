package application

import (
	"api/src/association/domain"
	"api/src/association/domain/ports"
)

type UpdateAssociationUseCase struct {
	repo ports.AssociationRepository
}

func NewUpdateAssociationUseCase(repo ports.AssociationRepository) *UpdateAssociationUseCase {
	return &UpdateAssociationUseCase{repo: repo}
}

func (useCase *UpdateAssociationUseCase) Execute(association domain.Association) error {
	return useCase.repo.Update(association)
}
