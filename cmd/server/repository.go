package main

import (
	"context"
	"database/sql"
	"time"
)

type CurrencyExchangeRepository struct {
	db *sql.DB
}

func NewCurrencyExchangeRepository(db *sql.DB) CurrencyExchangeRepository {
	return CurrencyExchangeRepository{
		db: db,
	}
}

func (r *CurrencyExchangeRepository) InsertCurrencyExchange(ctx context.Context, usdbrl USDBRL) error {
	statement, _ := r.db.Prepare("INSERT INTO currency (name, code, codein, bid, timestamp) VALUES (?, ?, ?, ?, ?)")
	
	_, err := statement.ExecContext(ctx, usdbrl.Name, usdbrl.Code, usdbrl.Codein, usdbrl.Bid, time.Now())
	if err != nil {
		return err
	}

	return nil
}