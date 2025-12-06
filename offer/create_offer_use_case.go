package offer

import "context"

type writer interface {
	Save(ctx context.Context, offer *Offer) error
}

type CreateOfferUseCase struct {
	writer writer
}

func NewCreateOfferUseCase(w writer) *CreateOfferUseCase {
	return &CreateOfferUseCase{writer: w}
}

func (uc *CreateOfferUseCase) Execute(ctx context.Context, offer *Offer) error {
	return uc.writer.Save(ctx, offer)
}
