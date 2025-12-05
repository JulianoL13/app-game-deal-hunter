package offer

import "context"

type GetOffersByGameNameUseCase struct {
	reader Reader
}

func NewGetOffersByGameNameUseCase(reader Reader) *GetOffersByGameNameUseCase {
	return &GetOffersByGameNameUseCase{reader: reader}
}

func (uc *GetOffersByGameNameUseCase) Execute(ctx context.Context, name string) ([]*Offer, error) {
	return uc.reader.FindByGameName(ctx, name)
}
