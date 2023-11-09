package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}
	u.format()
	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("the name field is required")
	}

	if u.Nick == "" {
		return errors.New("the nick field is required")
	}

	if u.Email == "" {
		return errors.New("the email field is required")
	}

	if u.Password == "" {
		return errors.New("the password field is required")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
