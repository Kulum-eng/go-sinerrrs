package ports

import "api/src/association/domain"

type AssociationRepository interface {
	Create(association domain.Association) error
	GetByID(id int) (*domain.Association, error)
	GetAll() ([]domain.Association, error)
	Update(association domain.Association) error
	Delete(id int) error
}
