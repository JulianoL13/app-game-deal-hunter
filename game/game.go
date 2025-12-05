package game

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrEmptyTitle = errors.New("game title cannot be empty")

type Game struct {
	ID          uuid.UUID
	Title       string
	Description string
	Developer   string
	Publisher   string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(title, description, developer, publisher string, releaseDate time.Time) (*Game, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Game{
		ID:          id,
		Title:       title,
		Description: description,
		Developer:   developer,
		Publisher:   publisher,
		ReleaseDate: releaseDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
