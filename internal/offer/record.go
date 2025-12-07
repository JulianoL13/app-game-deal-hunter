package offer

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	id         uuid.UUID
	offerID    uuid.UUID
	price      Price
	recordedAt time.Time
}

func NewRecord(offerID uuid.UUID, price Price) (*Record, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Record{
		id:         id,
		offerID:    offerID,
		price:      price,
		recordedAt: time.Now(),
	}, nil
}

func ReconstructRecord(id, offerID uuid.UUID, price Price, recordedAt time.Time) *Record {
	return &Record{
		id:         id,
		offerID:    offerID,
		price:      price,
		recordedAt: recordedAt,
	}
}

func (r *Record) ID() uuid.UUID         { return r.id }
func (r *Record) OfferID() uuid.UUID    { return r.offerID }
func (r *Record) Price() Price          { return r.price }
func (r *Record) RecordedAt() time.Time { return r.recordedAt }
