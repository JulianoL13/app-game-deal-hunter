package game

import (
	"context"
	"time"
)

type GameInput struct {
	Title       string
	Description string
	Developer   string
	Publisher   string
	ReleaseDate time.Time
}

type CreateGameUseCase struct {
	repo Repository
}

func NewCreateGameUseCase(repo Repository) *CreateGameUseCase {
	return &CreateGameUseCase{repo: repo}
}

func (uc *CreateGameUseCase) Execute(ctx context.Context, input GameInput) error {
	return uc.repo.Insert(ctx, &Game{
		Title:       input.Title,
		Description: input.Description,
		Developer:   input.Developer,
		Publisher:   input.Publisher,
		ReleaseDate: input.ReleaseDate,
	})
}
