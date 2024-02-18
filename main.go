package main

import (
	"context"
	"fmt"

	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/rand-quote-grpc"
	"github.com/mashmorsik/rand-quote-grpc/internal"
)

func main() {
	character := "Hagrid"
	quote, err := internal.ReturnQuote(character)
	if err != nil {
		fmt.Errorf("can't find quote for character: %s", character)
	}
	fmt.Println(quote)

	qs := QuotesServer{}

	grpc.
		randquotev1.RegisterRandQuotesServer()
}

type QuotesServer struct {
}

func (q QuotesServer) GetQuote(context.Context, *randquotev1.RandQuoteRequest,
) (*randquotev1.RandQuoteResponse, error) {
	return nil, nil
}
