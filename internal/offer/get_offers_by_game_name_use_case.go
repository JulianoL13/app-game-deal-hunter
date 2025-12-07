package offer

import "context"

type readerByGameName interface {
	FindByGameName(ctx context.Context, name string) ([]*Offer, error)
}

type GetOffersByGameNameUseCase struct {
	reader readerByGameName
}

func NewGetOffersByGameNameUseCase(r readerByGameName) *GetOffersByGameNameUseCase {
	return &GetOffersByGameNameUseCase{reader: r}
}

func (uc *GetOffersByGameNameUseCase) Execute(ctx context.Context, name string) ([]*Offer, error) {
	return uc.reader.FindByGameName(ctx, name)
}
