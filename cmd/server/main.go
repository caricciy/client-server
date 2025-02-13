package main

import (
	"client-server/internal/util"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	extApiUrl, ok := os.LookupEnv("EXT_API_URL")
	if !ok {
		log.Fatal("EXT_API_URL is required")
	}

	database := util.NewSQLite3Connection()
	defer database.Close()

	repository := NewCurrencyExchangeRepository(database)
	service := NewCurrencyExchangeService(extApiUrl, repository)

	http.HandleFunc("/", HandleGetCurrencyExchange(service))

	port := util.GetEnvStr("SERVER_PORT", ":8080")
	log.Printf("Server started at %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
