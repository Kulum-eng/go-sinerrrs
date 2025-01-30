package ports

import "api/src/membership/domain"

type MembershipRepository interface {
	CreateMembership(membership domain.Membership) error
	GetMembershipByID(id int) (domain.Membership, error)
	GetAllMemberships() ([]domain.Membership, error)
	UpdateMembership(membership domain.Membership) error
	DeleteMembership(id int) error
}
