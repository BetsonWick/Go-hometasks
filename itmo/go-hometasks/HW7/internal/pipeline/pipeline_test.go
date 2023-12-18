package pipeline

import (
	"context"
	"hw7/internal/generator"
	"hw7/internal/model"
	"testing"
	"time"
)

const ordersCount = 10

func TestNewOrderPipelineImplementation(t *testing.T) {
	obj := NewOrderPipelineImplementation()
	ctx := context.TODO()
	ordersArray := make([]model.OrderInitialized, 0, ordersCount)
	for i := 0; i < ordersCount; i++ {
		ordersArray = append(ordersArray, model.CreateEmptyOrder(i))
	}
	orders := generator.NewOrderGeneratorImplementation().GenerateOrdersStream(ctx, ordersArray)

	queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	proccessed := make(chan model.OrderProcessFinished)
	defer func() {
		cancel()
		close(proccessed)
	}()

	obj.Start(
		queryCtx,
		model.CreateEmptyActions(),
		orders,
		proccessed,
	)

	proc := make([]bool, ordersCount)

	for it := 0; it < ordersCount; it++ {
		var ord model.OrderProcessFinished
		select {
		case <-queryCtx.Done():
			t.Error("Timeout exceeded!")
			return
		case ord = <-proccessed:
		}
		num := ord.OrderFinishedExternalInteraction.OrderProcessStarted.OrderInitialized.OrderID
		if num < 0 || num > ordersCount {
			t.Errorf("Error number of order: %d", num)
		}
		if proc[num] {
			t.Errorf("Number of order repeated: %d", num)
		}
		proc[num] = true
	}
}
