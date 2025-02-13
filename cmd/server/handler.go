package main

import (
	"encoding/json"
	"net/http"
)

type ResponseBid struct {
	Bid string `json:"bid"`
}

func HandleGetCurrencyExchange(service CurrencyExchangeService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Content-Type", "application/json")

		currency, err := service.GetCurrencyExchange(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		dto := ResponseBid{
			Bid: currency.USDBRL.Bid,
		}

		json.NewEncoder(w).Encode(dto)
	}
}