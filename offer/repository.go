package offer

import (
	"context"

	"github.com/google/uuid"
)

type Writer interface {
	Insert(ctx context.Context, offer *Offer) error
}

type Reader interface {
	FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*Offer, error)
	FindByGameName(ctx context.Context, name string) ([]*Offer, error)
}
