package repositories

import (
	"database/sql"
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
