package models

type Password struct {
	Actual string `json:"actual"`
	New    string `json:"new"`
}
