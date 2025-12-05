package offer

import (
	"context"

	"github.com/google/uuid"
)

type GetOffersByGameIdUseCase struct {
	reader Reader
}

func NewGetOffersByGameIdUseCase(reader Reader) *GetOffersByGameIdUseCase {
	return &GetOffersByGameIdUseCase{reader: reader}
}

func (uc *GetOffersByGameIdUseCase) Execute(ctx context.Context, gameID uuid.UUID) ([]*Offer, error) {
	return uc.reader.FindByGameID(ctx, gameID)
}
