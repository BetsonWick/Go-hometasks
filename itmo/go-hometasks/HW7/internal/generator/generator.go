package generator

import (
	"context"

	"hw7/internal/model"
)

const chanCount = 10

type OrderGenerator interface {
	GenerateOrdersStream(ctx context.Context, orders []model.OrderInitialized) <-chan model.OrderInitialized
}

type OrderGeneratorImplementation struct{}

func NewOrderGeneratorImplementation() *OrderGeneratorImplementation {
	return &OrderGeneratorImplementation{}
}

func (o *OrderGeneratorImplementation) GenerateOrdersStream(ctx context.Context, orders []model.OrderInitialized) <-chan model.OrderInitialized {
	ch := make(chan model.OrderInitialized, chanCount)

	go func() {
		defer close(ch)
		for _, order := range orders {
			select {
			case <-ctx.Done():
				return
			case ch <- order:
			}
		}
	}()

	return ch
}
