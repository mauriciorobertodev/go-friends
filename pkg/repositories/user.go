package repositories

import (
	"database/sql"
	"fmt"
	"go-friends/pkg/models"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (r users) Store(user models.User) (uint64, error) {
	stm, err := r.db.Prepare("INSERT INTO users (name, nick, email, password) values (?, ?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	defer stm.Close()

	result, err := stm.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, nil
	}

	return uint64(lastId), nil
}

func (r users) Search(search string) ([]models.User, error) {
	search = fmt.Sprintf("%%%s%%", search)

	rows, err := r.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?",
		search,
		search,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
