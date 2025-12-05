package offer

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidStore    = errors.New("invalid store")
	ErrInvalidPlatform = errors.New("invalid platform")
	ErrEmptyURL        = errors.New("offer URL cannot be empty")
)

type Offer struct {
	ID           uuid.UUID
	GameID       uuid.UUID
	Store        Store
	Platform     Platform
	URL          string
	CurrentPrice Price
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func New(gameID uuid.UUID, store Store, platform Platform, url string) (*Offer, error) {
	if !store.IsValid() {
		return nil, ErrInvalidStore
	}
	if !platform.IsValid() {
		return nil, ErrInvalidPlatform
	}
	if url == "" {
		return nil, ErrEmptyURL
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Offer{
		ID:        id,
		GameID:    gameID,
		Store:     store,
		Platform:  platform,
		URL:       url,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (o *Offer) UpdatePrice(price Price) {
	o.CurrentPrice = price
	o.UpdatedAt = time.Now()
}
