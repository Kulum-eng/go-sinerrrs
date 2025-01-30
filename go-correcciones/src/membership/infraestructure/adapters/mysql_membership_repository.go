package adapters

import (
	"api/src/membership/domain"
	"api/src/membership/domain/ports"
	"database/sql"
	"errors"
)

type MySQLMembershipRepository struct {
	db *sql.DB
}

func NewMySQLMembershipRepository(db *sql.DB) ports.MembershipRepository {
	return &MySQLMembershipRepository{db: db}
}

func (r *MySQLMembershipRepository) CreateMembership(m domain.Membership) error {
	query := "INSERT INTO memberships (user_id, association_id, status, role) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, m.UserID, m.AssociationID, m.Status, m.Role)
	return err
}
func (r *MySQLMembershipRepository) GetMembershipByID(id int) (domain.Membership, error) {
	var m domain.Membership
	query := "SELECT id, user_id, association_id, status, role FROM memberships WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&m.ID, &m.UserID, &m.AssociationID, &m.Status, &m.Role)
	if err != nil {
		return domain.Membership{}, errors.New("membership not found")
	}
	return m, nil
}

func (r *MySQLMembershipRepository) GetAllMemberships() ([]domain.Membership, error) {
	query := "SELECT id, user_id, association_id, status, role FROM memberships"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memberships []domain.Membership
	for rows.Next() {
		var m domain.Membership
		if err := rows.Scan(&m.ID, &m.UserID, &m.AssociationID, &m.Status, &m.Role); err != nil {
			return nil, err
		}
		memberships = append(memberships, m)
	}
	return memberships, nil
}

func (r *MySQLMembershipRepository) UpdateMembership(m domain.Membership) error {
	query := "UPDATE memberships SET status = ?, role = ? WHERE id = ?"
	_, err := r.db.Exec(query, m.Status, m.Role, m.ID)
	return err
}

func (r *MySQLMembershipRepository) DeleteMembership(id int) error {
	query := "DELETE FROM memberships WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
