package statisticooddscompiler

import (
	"context"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OddCompilerClient interface {
	GetEventMarket(ctx context.Context, eventID uint64, market string) (*statistico.EventMarket, error)
}

type oddCompilerClient struct {
	client statistico.OddsCompilerServiceClient
}

func (o *oddCompilerClient) GetEventMarket(ctx context.Context, eventID uint64, market string) (*statistico.EventMarket, error) {
	req := statistico.EventRequest{
		EventId: eventID,
		Market:  market,
	}

	em, err := o.client.GetEventMarket(ctx, &req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return nil, ErrorNotFound{eventID, err}
			case codes.FailedPrecondition:
				return nil, ErrorInvalidArgument{err}
			default:
				return nil, ErrorBadGateway{err}
			}
		}
	}

	return em, nil
}

func NewOddsCompilerClient(c statistico.OddsCompilerServiceClient) OddCompilerClient {
	return &oddCompilerClient{client: c}
}
