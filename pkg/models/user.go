package models

import "time"

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
