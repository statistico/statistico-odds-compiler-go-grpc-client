package statisticooddscompiler_test

import (
	"context"
	"github.com/statistico/statistico-odds-compiler-go-grpc-client"
	"github.com/statistico/statistico-proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestOddsCompilerClient_GetEventMarket(t *testing.T) {
	t.Run("calls odd compiler service and returns an EventMarket struct", func(t *testing.T) {
		t.Helper()

		mc := new(MockProtoOddsCompilerClient)
		client := statisticooddscompiler.NewOddsCompilerClient(mc)

		ctx := context.Background()

		r := mock.MatchedBy(func(r *statistico.EventRequest) bool {
			assert.Equal(t, uint64(564), r.GetEventId())
			assert.Equal(t, "OVER_UNDER_25", r.GetMarket())
			return true
		})

		market := statistico.EventMarket{
			EventId: 564,
			Market:  "OVER_UNDER_25",
			Odds: []*statistico.Odds{
				{
					Price:     1.95,
					Selection: "Over",
				},
				{
					Price:     1.95,
					Selection: "under",
				},
			},
		}

		mc.On("GetEventMarket", ctx, r, []grpc.CallOption(nil)).Return(&market, nil)

		fetched, err := client.GetEventMarket(ctx, uint64(564), "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, fetched, &market)
		mc.AssertExpectations(t)
	})

	t.Run("returns an error if failed precondition error returned by service", func(t *testing.T) {
		t.Helper()

		mc := new(MockProtoOddsCompilerClient)
		client := statisticooddscompiler.NewOddsCompilerClient(mc)

		ctx := context.Background()

		r := mock.MatchedBy(func(r *statistico.EventRequest) bool {
			assert.Equal(t, uint64(564), r.GetEventId())
			assert.Equal(t, "OVER_UNDER_25", r.GetMarket())
			return true
		})

		e := status.Error(codes.FailedPrecondition, "failed precondition")

		mc.On("GetEventMarket", ctx, r, []grpc.CallOption(nil)).Return(&statistico.EventMarket{}, e)

		_, err := client.GetEventMarket(ctx, uint64(564), "OVER_UNDER_25")

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "invalid argument provided: rpc error: code = FailedPrecondition desc = failed precondition", err.Error())
		mc.AssertExpectations(t)
	})

	t.Run("returns an error if not found error returned by service", func(t *testing.T) {
		t.Helper()

		mc := new(MockProtoOddsCompilerClient)
		client := statisticooddscompiler.NewOddsCompilerClient(mc)

		ctx := context.Background()

		r := mock.MatchedBy(func(r *statistico.EventRequest) bool {
			assert.Equal(t, uint64(564), r.GetEventId())
			assert.Equal(t, "OVER_UNDER_25", r.GetMarket())
			return true
		})

		e := status.Error(codes.NotFound, "not found")

		mc.On("GetEventMarket", ctx, r, []grpc.CallOption(nil)).Return(&statistico.EventMarket{}, e)

		_, err := client.GetEventMarket(ctx, uint64(564), "OVER_UNDER_25")

		if err == nil {
			t.Fatal("Expected errors, got nil")
		}

		assert.Equal(t, "resource with is '564' does not exist. Error: rpc error: code = NotFound desc = not found", err.Error())
		mc.AssertExpectations(t)
	})
}

type MockProtoOddsCompilerClient struct {
	mock.Mock
}

func (m *MockProtoOddsCompilerClient) GetEventMarket(ctx context.Context, in *statistico.EventRequest, opts ...grpc.CallOption) (*statistico.EventMarket, error) {
	args := m.Called(ctx, in, opts)
	return args.Get(0).(*statistico.EventMarket), args.Error(1)
}
