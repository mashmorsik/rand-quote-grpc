package handlers

import (
	"context"
	"fmt"
	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/rand-quote-grpc"
	"github.com/mashmorsik/rand-quote-grpc/internal/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QuotesServer struct {
	*randquotev1.UnimplementedRandQuotesServer
}

func (q QuotesServer) GetQuote(ctx context.Context, req *randquotev1.RandQuoteRequest,
) (*randquotev1.RandQuoteResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("name of the character is required")
	}

	name := req.GetName()
	quote, err := pkg.ReturnQuote(name.String())
	if err != nil {
		return nil, fmt.Errorf("can't find quote by %s", name)
	}

	return &randquotev1.RandQuoteResponse{Quote: quote}, nil
}

func (q QuotesServer) ListQuotes(req *randquotev1.ListQuotesRequest, stream randquotev1.RandQuotes_ListQuotesServer,
) error {

	if req == nil {
		return fmt.Errorf("name of the character is required")
	}

	name := req.GetName()
	ctx := stream.Context()

	quotes, err := pkg.ReturnQuotesList(name.String())
	if err != nil {
		fmt.Println(fmt.Errorf("can't find quote by %s", name))
		return fmt.Errorf("can't find quote by %s", name)
	}

	if len(quotes) == 0 {
		return status.Errorf(codes.NotFound, "can't find any quote by %s", name)
	}

	for i := 0; i < len(quotes); i++ {
		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			return ctx.Err()
		}

		err = stream.Send(&randquotev1.ListQuotesResponse{Quotes: quotes[i]})
		if err != nil {
			fmt.Println(fmt.Errorf("can't send message: %s to stream", quotes[i]))
			return fmt.Errorf("can't send message: %s to stream", quotes[i])
		}
	}

	err = stream.Send(&randquotev1.ListQuotesResponse{Quotes: "That's it."})
	if err != nil {
		return err
	}

	return nil
}
