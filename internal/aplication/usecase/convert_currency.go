package usecase

import (
	"errors"

	"github.com/marceloneiva/myapi/internal/domain/entity"
)

type ConvertCurrencyUseCase struct {
	rateRepo ExchangeRateRepository
}

type ExchangeRateRepository interface {
	GetRate(from, to entity.Currency) (*entity.ExchangeRate, error)
}

func NewConvertCurrencyUseCase(r ExchangeRateRepository) *ConvertCurrencyUseCase {
	return &ConvertCurrencyUseCase{rateRepo: r}
}

func (uc *ConvertCurrencyUseCase) Execute(from, to entity.Currency, amount float64) (float64, error) {
	rate, err := uc.rateRepo.GetRate(from, to)
	if err != nil {
		return 0, errors.New("não foi possível obter a taxa de câmbio")
	}
	return rate.Convert(amount), nil
}
