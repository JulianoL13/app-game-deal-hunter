package offer

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	ID         uuid.UUID
	OfferID    uuid.UUID
	Price      Price
	RecordedAt time.Time
}

func NewHistory(offerID uuid.UUID, price Price) (*Record, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Record{
		ID:         id,
		OfferID:    offerID,
		Price:      price,
		RecordedAt: time.Now(),
	}, nil
}
