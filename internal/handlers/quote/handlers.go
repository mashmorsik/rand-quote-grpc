package quote

import (
	"context"
	"fmt"
	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/rand-quote-grpc"
	"github.com/mashmorsik/rand-quote-grpc/internal/quote"
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
	quote, err := quote.ReturnQuote(name)
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

	quotes, err := quote.ReturnQuotesList(name)
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

		err = stream.Send(&randquotev1.ListQuotesResponse{Quote: quotes[i]})
		if err != nil {
			fmt.Println(fmt.Errorf("can't send message: %s to stream", quotes[i]))
			return fmt.Errorf("can't send message: %s to stream", quotes[i])
		}
	}

	err = stream.Send(&randquotev1.ListQuotesResponse{Quote: "That's it."})
	if err != nil {
		return err
	}

	return nil
}

func (q QuotesServer) GetSeveralCharactersQuotes(stream randquotev1.RandQuotes_GetSeveralCharactersQuotesServer) error {
	quotes := make(map[string]string)

	for i := 0; len(quotes) < 3; i++ {
		name, err := stream.Recv()
		if err != nil {
			fmt.Println(fmt.Errorf("can't receive message"))
			return fmt.Errorf("can't receive message")
		}

		fmt.Println(name)

		rQuote, err := quote.ReturnQuote(name.Name)
		if err != nil {
			fmt.Println(fmt.Errorf("can't find rQuote by %s", name))
			return fmt.Errorf("can't find rQuote by %s", name)
		}

		nameStr, err := quote.EnumMatcher(name.Name)
		if err != nil {
			fmt.Println(fmt.Errorf("can't convert enum: %s to string", name))
			return fmt.Errorf("can't convert enum: %s to string", name)
		}

		quotes[nameStr] = rQuote
	}
	return stream.SendAndClose(&randquotev1.SeveralCharacterQuotesResponse{Quotes: quotes})
}

func (q QuotesServer) QuotesChat(stream randquotev1.RandQuotes_QuotesChatServer) error {
	name, err := stream.Recv()
	if err != nil {
		fmt.Println(fmt.Errorf("can't receive message"))
		return fmt.Errorf("can't receive message")
	}
	if name.Name != randquotev1.Name_CLOSE_CONNECTION {
		quote, err := quote.ReturnQuote(name.Name)
		if err != nil {
			return nil
		}
		err = stream.Send(&randquotev1.QuotesChatResponse{Quote: quote})
		if err != nil {
			fmt.Println(fmt.Errorf("can't send message: %s to stream", quote))
			return fmt.Errorf("can't send message: %s to stream", quote)
		}

		err = q.QuotesChat(stream)
		if err != nil {
			fmt.Println(fmt.Errorf("can't continue stream"))
			return fmt.Errorf("can't continue stream")
		}
	}

	return nil
}
