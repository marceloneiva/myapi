package entity

import "errors"

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	BRL Currency = "BRL"
)

type ExchangeRate struct {
	From Currency
	To   Currency
	Rate float64
}

func NewExchangeRate(from Currency, to Currency, rate float64) (*ExchangeRate, error) {
	if rate <= 0 {
		return nil, errors.New("taxa de câmbio deve ser positiva")
	}
	return &ExchangeRate{From: from, To: to, Rate: rate}, nil
}

func (e *ExchangeRate) Convert(amount float64) float64 {
	return amount * e.Rate
}
