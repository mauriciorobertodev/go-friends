package repositories

import (
	"database/sql"
	"errors"
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

func (r users) FindUserById(id uint64) (models.User, error) {
	rows, err := r.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)

		if err != nil {
			return models.User{}, err
		}

		return user, nil
	}

	return models.User{}, errors.New("user does not exists")
}

func (r users) UpdateUser(id uint64, user models.User) error {
	stm, err := r.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(user.Name, user.Nick, user.Email, id)

	if err != nil {
		return err
	}

	return nil
}

func (r users) DeleteUser(id uint64) error {
	stm, err := r.db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
