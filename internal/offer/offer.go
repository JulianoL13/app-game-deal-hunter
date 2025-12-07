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
	ErrNilPrice   = errors.New("price cannot be nil")
)

type Offer struct {
	id           uuid.UUID
	gameID       uuid.UUID
	store        Store
	url          string
	currentPrice Price
	createdAt    time.Time
	updatedAt    time.Time
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
		id:        id,
		gameID:    gameID,
		store:     store,
		url:       rawURL,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func ReconstructOffer(id, gameID uuid.UUID, store Store, url string, currentPrice Price, createdAt, updatedAt time.Time) *Offer {
	return &Offer{
		id:           id,
		gameID:       gameID,
		store:        store,
		url:          url,
		currentPrice: currentPrice,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

func (o *Offer) UpdatePrice(price Price) {
	o.currentPrice = price
	o.updatedAt = time.Now()
}

func (o *Offer) ChangeURL(rawURL string) error {
	if rawURL == "" {
		return ErrEmptyURL
	}

	u, err := url.Parse(rawURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return ErrInvalidURL
	}

	o.url = rawURL
	o.updatedAt = time.Now()
	return nil
}

func (o *Offer) ID() uuid.UUID        { return o.id }
func (o *Offer) GameID() uuid.UUID    { return o.gameID }
func (o *Offer) Store() Store         { return o.store }
func (o *Offer) URL() string          { return o.url }
func (o *Offer) CurrentPrice() Price  { return o.currentPrice }
func (o *Offer) CreatedAt() time.Time { return o.createdAt }
func (o *Offer) UpdatedAt() time.Time { return o.updatedAt }

func (o *Offer) HasPrice() bool {
	return !o.currentPrice.IsZero()
}

func (o *Offer) IsFree() bool {
	return o.currentPrice.Amount() == 0 && o.HasPrice()
}
