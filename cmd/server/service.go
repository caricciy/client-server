package main

import (
	"client-server/internal/util"
	"context"
	"errors"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

type USDBRL struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type CurrencyExchange struct {
	USDBRL USDBRL `json:"USDBRL"`
}

type CurrencyExchangeService struct{
	exchangeApiUrl string
	repository CurrencyExchangeRepository
}

func NewCurrencyExchangeService(exchangeApiUrl string, repository CurrencyExchangeRepository) CurrencyExchangeService {
	return CurrencyExchangeService{
		exchangeApiUrl: exchangeApiUrl,
		repository: repository,
	}
}

func (s *CurrencyExchangeService) GetCurrencyExchange(ctx context.Context) (*CurrencyExchange, error) {
	opts := util.GetOpts{
		Timeout: 200 * time.Millisecond,
		URL:     s.exchangeApiUrl,
	}

	var currency CurrencyExchange
	if err := util.GetWithTimeout(ctx, &currency, opts); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("failed to get currency exchange: %v", err)
		}
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10 * time.Millisecond)
	defer cancel()

	if err := s.repository.InsertCurrencyExchange(ctx, currency.USDBRL); err != nil {	
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("failed to insert currency exchange: %v", err)
		}
	}

	return &currency, nil
}