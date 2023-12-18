package workerpool

import (
	"context"
	"hw7/internal/generator"
	"hw7/internal/model"
	"testing"
	"time"
)

const ORDERS_AMOUNT = 10

func TestNewOrderPipelineImplementation(t *testing.T) {
	obj := NewOrderWorkerPoolImplementation()
	ctx := context.TODO()
	ordersArray := make([]model.OrderInitialized, 0, ORDERS_AMOUNT)
	for i := 0; i < ORDERS_AMOUNT; i++ {
		ordersArray = append(ordersArray, model.CreateEmptyOrder(i))
	}
	orders := generator.NewOrderGeneratorImplementation().GenerateOrdersStream(ctx, ordersArray)

	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	proccessed := obj.StartWorkerPool(
		queryCtx,
		orders,
		model.CreateEmptyActions(),
		10,
	)

	defer cancel()

	proc := make([]bool, ORDERS_AMOUNT)

	for it := 0; it < ORDERS_AMOUNT; it++ {
		var ord model.OrderProcessFinished
		select {
		case <-queryCtx.Done():
			t.Error("Timeout exceeded!")
			return
		case ord = <-proccessed:
		}
		num := ord.OrderFinishedExternalInteraction.OrderProcessStarted.OrderInitialized.OrderID
		if num < 0 || num > ORDERS_AMOUNT {
			t.Errorf("Error number of order: %d", num)
		}
		if proc[num] {
			t.Errorf("Number of order repeated: %d", num)
		}
		proc[num] = true
	}
}
