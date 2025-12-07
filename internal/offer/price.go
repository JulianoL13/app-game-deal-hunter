package offer

import "time"

type Price struct {
	amount    int64
	currency  string
	updatedAt time.Time
}

func NewPrice(amount int64, currency string) Price {
	return Price{
		amount:    amount,
		currency:  currency,
		updatedAt: time.Now(),
	}
}

func ReconstructPrice(amount int64, currency string, updatedAt time.Time) Price {
	return Price{
		amount:    amount,
		currency:  currency,
		updatedAt: updatedAt,
	}
}

func (p Price) Amount() int64        { return p.amount }
func (p Price) Currency() string     { return p.currency }
func (p Price) UpdatedAt() time.Time { return p.updatedAt }

func (p Price) Reais() float64 {
	return float64(p.amount) / 100
}

func (p Price) IsZero() bool {
	return p.amount == 0 && p.currency == ""
}

func (p Price) Equals(other Price) bool {
	return p.amount == other.amount && p.currency == other.currency
}

func (p Price) WithAmount(amount int64) Price {
	return Price{
		amount:    amount,
		currency:  p.currency,
		updatedAt: time.Now(),
	}
}

func (p Price) WithCurrency(currency string) Price {
	return Price{
		amount:    p.amount,
		currency:  currency,
		updatedAt: time.Now(),
	}
}
