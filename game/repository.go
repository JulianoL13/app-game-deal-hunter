package game

import "context"

type Writer interface {
	Insert(ctx context.Context, game *Game) error
}

type Repository interface {
	Writer
	Close() error
}
