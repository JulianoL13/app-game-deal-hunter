package game

import "context"

type writer interface {
	Save(ctx context.Context, game *Game) error
}

type CreateGameUseCase struct {
	writer writer
}

func NewCreateGameUseCase(w writer) *CreateGameUseCase {
	return &CreateGameUseCase{writer: w}
}

func (uc *CreateGameUseCase) Execute(ctx context.Context, game *Game) error {
	return uc.writer.Save(ctx, game)
}
