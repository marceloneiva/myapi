package repository

import "github.com/marceloneiva/myapi/internal/domain/entity"

type InMemoryRateRepo struct {
	rates map[string]*entity.ExchangeRate
}

func NewInMemoryRateRepo() *InMemoryRateRepo {
	return &InMemoryRateRepo{
		rates: map[string]*entity.ExchangeRate{
			"USD-BRL": {From: entity.USD, To: entity.BRL, Rate: 5.10},
			"BRL-USD": {From: entity.BRL, To: entity.USD, Rate: 0.20},
		},
	}
}

func (r *InMemoryRateRepo) GetRate(from, to entity.Currency) (*entity.ExchangeRate, error) {
	key := string(from) + "-" + string(to)
	if rate, ok := r.rates[key]; ok {
		return rate, nil
	}
	return nil, nil
}
