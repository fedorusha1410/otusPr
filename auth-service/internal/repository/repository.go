package repository

import (
	"auth-service/internal/model/user"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUsers() ([]*user.User, error) {
	rows, err := r.db.Query("SELECT id, name, role, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		u := &user.User{}
		var roleInt int
		if err := rows.Scan(&u.Id, &u.Name, &roleInt, &u.Password); err != nil {
			return nil, err
		}
		u.Role = user.Role(roleInt)
		users = append(users, u)
	}
	return users, nil
}

func (r *Repository) GetUserById(id int) (*user.User, error) {

	query := "SELECT id, name, role, password FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	u := &user.User{}
	var roleInt int
	err := row.Scan(&u.Id, &u.Name, &roleInt, &u.Password)
	if err != nil {
		return nil, err
	}
	u.Role = user.Role(roleInt)
	return u, nil
}

func (r *Repository) GetByUsername(username string) (*user.User, error) {
	query := `SELECT id, name, role, password FROM users WHERE name = $1`
	row := r.db.QueryRow(query, username)

	u := &user.User{}
	var roleInt int
	err := row.Scan(&u.Id, &u.Name, &roleInt, &u.Password)
	if err != nil {
		return nil, err
	}

	u.Role = user.Role(roleInt)
	return u, nil
}

func (r *Repository) UpdateUser(id int, newData *user.User) error {

	query := `UPDATE users SET name = $1, role = $2, password = $3 WHERE id = $4`
	_, err := r.db.Exec(query, newData.Name, int(newData.Role), newData.Password, id)
	return err
}

func (r *Repository) DeleteUser(id int) error {

	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) Save(newUser *user.User) error {
	query := `INSERT INTO users (name, role, password) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, newUser.Name, int(newUser.Role), newUser.Password).Scan(&newUser.Id)
}
