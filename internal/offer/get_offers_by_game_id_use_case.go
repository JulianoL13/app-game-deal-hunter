package offer

import (
	"context"

	"github.com/google/uuid"
)

type readerByGameID interface {
	FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*Offer, error)
}

type GetOffersByGameIdUseCase struct {
	reader readerByGameID
}

func NewGetOffersByGameIdUseCase(r readerByGameID) *GetOffersByGameIdUseCase {
	return &GetOffersByGameIdUseCase{reader: r}
}

func (uc *GetOffersByGameIdUseCase) Execute(ctx context.Context, gameID uuid.UUID) ([]*Offer, error) {
	return uc.reader.FindByGameID(ctx, gameID)
}
