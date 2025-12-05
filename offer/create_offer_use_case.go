package offer

import "context"

type CreateOfferUseCase struct {
	writer Writer
}

func NewCreateOfferUseCase(writer Writer) *CreateOfferUseCase {
	return &CreateOfferUseCase{writer: writer}
}

func (uc *CreateOfferUseCase) Execute(ctx context.Context, offer *Offer) error {
	return uc.writer.Insert(ctx, offer)
}
