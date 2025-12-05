package offer

import "time"

type Price struct {
	Amount    int64
	Currency  string
	UpdatedAt time.Time
}

func NewPrice(amount int64, currency string) Price {
	return Price{
		Amount:    amount,
		Currency:  currency,
		UpdatedAt: time.Now(),
	}
}

func (p Price) Reais() float64 {
	return float64(p.Amount) / 100
}

func (p Price) IsZero() bool {
	return p.Amount == 0 && p.Currency == ""
}

func (p Price) Equals(other Price) bool {
	return p.Amount == other.Amount && p.Currency == other.Currency
}
