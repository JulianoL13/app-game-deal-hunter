package game

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyTitle     = errors.New("game title cannot be empty")
	ErrEmptyDeveloper = errors.New("game developer cannot be empty")
)

type Game struct {
	id          uuid.UUID
	title       string
	description string
	developer   string
	publisher   string
	releaseDate time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func NewGame(title, description, developer, publisher string, releaseDate time.Time) (*Game, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Game{
		id:          id,
		title:       title,
		description: description,
		developer:   developer,
		publisher:   publisher,
		releaseDate: releaseDate,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

func Reconstruct(id uuid.UUID, title, description, developer, publisher string, releaseDate, createdAt, updatedAt time.Time) *Game {
	return &Game{
		id:          id,
		title:       title,
		description: description,
		developer:   developer,
		publisher:   publisher,
		releaseDate: releaseDate,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (g *Game) Rename(newTitle string) error {
	if newTitle == "" {
		return ErrEmptyTitle
	}
	g.title = newTitle
	g.updatedAt = time.Now()
	return nil
}

func (g *Game) UpdateDescription(description string) {
	g.description = description
	g.updatedAt = time.Now()
}

func (g *Game) ChangeDeveloper(developer string) error {
	if developer == "" {
		return ErrEmptyDeveloper
	}
	g.developer = developer
	g.updatedAt = time.Now()
	return nil
}

func (g *Game) ChangePublisher(publisher string) {
	g.publisher = publisher
	g.updatedAt = time.Now()
}

func (g *Game) ID() uuid.UUID          { return g.id }
func (g *Game) Title() string          { return g.title }
func (g *Game) Description() string    { return g.description }
func (g *Game) Developer() string      { return g.developer }
func (g *Game) Publisher() string      { return g.publisher }
func (g *Game) ReleaseDate() time.Time { return g.releaseDate }
func (g *Game) CreatedAt() time.Time   { return g.createdAt }
func (g *Game) UpdatedAt() time.Time   { return g.updatedAt }

func (g *Game) IsReleased() bool {
	return !g.releaseDate.IsZero() && g.releaseDate.Before(time.Now())
}
