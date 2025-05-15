package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marceloneiva/myapi/internal/config"
	"github.com/marceloneiva/myapi/internal/domain/entity"
)

type MySQLRateRepo struct {
	db *sql.DB
}

func NewMySQLRateRepo() *MySQLRateRepo {
	return &MySQLRateRepo{db: config.DB}
}

func (r *MySQLRateRepo) GetRate(from, to entity.Currency) (*entity.ExchangeRate, error) {
	query := `SELECT rate FROM exchange_rates WHERE from_currency = ? AND to_currency = ? LIMIT 1`
	row := r.db.QueryRowContext(context.Background(), query, from, to)

	var rate float64
	if err := row.Scan(&rate); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("taxa de câmbio não encontrada")
		}
		return nil, err
	}

	return entity.NewExchangeRate(from, to, rate)
}
