package main

import (
	"client-server/internal/util"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

type ResponseDataBid struct {
	Bid string `json:"bid"`
}

func (r *ResponseDataBid) String() string {
	return fmt.Sprintf("Dólar:{%s}\n", r.Bid)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	opts := util.GetOpts{
		Timeout: 300 * time.Millisecond,
		URL:     os.Getenv("SERVER_URL"),
	}
	
	var resp ResponseDataBid
	if err := util.GetWithTimeout(context.Background(), &resp, opts); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Fatalf("failed calling currency exchange api: %v", err)
		}
	}

	filename := fmt.Sprintf("%s/cotacao.txt", os.Getenv("OUTPUT_DIR"))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = "./cotacao.txt"
	}

	if err := util.WriteToFile(filename, resp.String()); err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}