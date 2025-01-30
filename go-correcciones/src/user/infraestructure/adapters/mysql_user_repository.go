package adapters

import (
	"api/src/user/domain"
	"database/sql"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{DB: db}
}

func (repo *MySQLUserRepository) CreateUser(user domain.User) error {
	_, err := repo.DB.Exec("INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Password, user.Role)
	return err
}

func (repo *MySQLUserRepository) GetUserByID(id int) (domain.User, error) {
	var user domain.User
	err := repo.DB.QueryRow("SELECT id, name, email, password, role FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	return user, err
}

func (repo *MySQLUserRepository) GetAllUsers() ([]domain.User, error) {
	rows, err := repo.DB.Query("SELECT id, name, email, password, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (repo *MySQLUserRepository) UpdateUser(user domain.User) error {
	_, err := repo.DB.Exec("UPDATE users SET name=?, email=?, password=?, role=? WHERE id=?",
		user.Name, user.Email, user.Password, user.Role, user.ID)
	return err
}

func (repo *MySQLUserRepository) DeleteUser(id int) error {
	_, err := repo.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
