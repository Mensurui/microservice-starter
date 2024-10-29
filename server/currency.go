package server

import (
	"context"

	protos "github.com/Mensurui/microservice-starter/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{
		UnimplementedCurrencyServer: protos.UnimplementedCurrencyServer{}, // Correctly initialize
		log:                         l,
	}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle get rate", "base", rr.GetBase(), "destination", rr.GetDestination)
	return &protos.RateResponse{
		Rate: 0.5,
	}, nil
}
