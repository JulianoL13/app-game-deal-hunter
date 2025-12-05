package pricehistory

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	ID         uuid.UUID
	OfferID    uuid.UUID
	Amount     int64
	Currency   string
	RecordedAt time.Time
}

func New(offerID uuid.UUID, amount int64, currency string) (*Record, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Record{
		ID:         id,
		OfferID:    offerID,
		Amount:     amount,
		Currency:   currency,
		RecordedAt: time.Now(),
	}, nil
}
