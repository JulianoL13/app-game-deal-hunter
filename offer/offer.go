package offer

import (
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyURL   = errors.New("offer URL cannot be empty")
	ErrInvalidURL = errors.New("offer URL must be a valid HTTP/HTTPS URL")
)

type Offer struct {
	ID           uuid.UUID
	GameID       uuid.UUID
	Store        Store
	URL          string
	CurrentPrice Price
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewOffer(gameID uuid.UUID, store Store, rawURL string) (*Offer, error) {
	if rawURL == "" {
		return nil, ErrEmptyURL
	}

	u, err := url.Parse(rawURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return nil, ErrInvalidURL
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
		URL:       rawURL,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (o *Offer) UpdatePrice(price Price) {
	o.CurrentPrice = price
	o.UpdatedAt = time.Now()
}
