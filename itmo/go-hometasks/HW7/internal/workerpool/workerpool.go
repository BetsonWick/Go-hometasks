package workerpool

import (
	"context"
	"hw7/internal/pipeline"
	"sync"

	"hw7/internal/model"
)

type OrderWorkerPool interface {
	StartWorkerPool(ctx context.Context, orders <-chan model.OrderInitialized, additionalActions model.OrderActions, workersCount int) <-chan model.OrderProcessFinished
}

type OrderWorkerPoolImplementation struct{}

func NewOrderWorkerPoolImplementation() *OrderWorkerPoolImplementation {
	return &OrderWorkerPoolImplementation{}
}

func (o *OrderWorkerPoolImplementation) StartWorkerPool(ctx context.Context, orders <-chan model.OrderInitialized, additionalActions model.OrderActions, workersCount int) <-chan model.OrderProcessFinished {
	var wg sync.WaitGroup
	results := make(chan model.OrderProcessFinished, workersCount)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pipeline.NewOrderPipelineImplementation().Start(ctx, additionalActions, orders, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
