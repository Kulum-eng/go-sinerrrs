package application

import (
	"api/src/membership/domain/ports"
)

type DeleteMembershipUseCase struct {
	repo ports.MembershipRepository
}

func NewDeleteMembershipUseCase(repo ports.MembershipRepository) *DeleteMembershipUseCase {
	return &DeleteMembershipUseCase{repo: repo}
}

func (uc *DeleteMembershipUseCase) Execute(id int) error {
	return uc.repo.DeleteMembership(id)
}
