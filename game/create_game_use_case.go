package game

import "context"

type CreateGameUseCase struct {
	writer Writer
}

func NewCreateGameUseCase(writer Writer) *CreateGameUseCase {
	return &CreateGameUseCase{writer: writer}
}

func (uc *CreateGameUseCase) Execute(ctx context.Context, game *Game) error {
	return uc.writer.Insert(ctx, game)
}
