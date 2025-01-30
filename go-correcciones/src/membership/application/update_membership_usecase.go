package application

import (
	"api/src/membership/domain"
	"api/src/membership/domain/ports"
)

type UpdateMembershipUseCase struct {
	repo ports.MembershipRepository
}
func NewUpdateMembershipUseCase(repo ports.MembershipRepository) *UpdateMembershipUseCase {
	return &UpdateMembershipUseCase{repo: repo}
}

func (uc *UpdateMembershipUseCase) Execute(membership domain.Membership) error {
	return uc.repo.UpdateMembership(membership)
}
