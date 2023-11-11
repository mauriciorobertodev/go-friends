package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"create_at,omitempty"`
}

func (p *Post) Prepare(register bool) error {
	if err := p.validate(); err != nil {
		return err
	}
	err := p.format()

	if err != nil {
		return err
	}

	return nil
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("the title field is required")
	}

	if p.Content == "" {
		return errors.New("the content field is required")
	}

	return nil
}

func (p *Post) format() error {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
	return nil
}
