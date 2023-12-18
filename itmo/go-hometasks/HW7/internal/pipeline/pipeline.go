package pipeline

import (
	"context"
	"hw7/internal/model"
)

const fanLimit = 10

type OrderPipeline interface {
	Start(ctx context.Context, actions model.OrderActions, orders <-chan model.OrderInitialized, processed chan<- model.OrderProcessFinished)
}

type OrderPipelineImplementation struct{}

func NewOrderPipelineImplementation() *OrderPipelineImplementation {
	return &OrderPipelineImplementation{}
}

func StartStage(
	ctx context.Context,
	actions *model.OrderActions,
	orders <-chan model.OrderInitialized,
) <-chan model.OrderProcessStarted {
	outCh := make(chan model.OrderProcessStarted, 10)
	go func() {
		defer close(outCh)
		for order := range orders {
			select {
			case <-ctx.Done():
				return
			case outCh <- order.NextActionState(actions):
			}
		}
	}()
	return outCh
}

func ProcessStage(
	ctx context.Context,
	actions *model.OrderActions,
	orders <-chan model.OrderProcessStarted,
) <-chan model.OrderFinishedExternalInteraction {
	outCh := make(chan model.OrderFinishedExternalInteraction, 10)
	go func() {
		defer close(outCh)
		for order := range orders {
			select {
			case <-ctx.Done():
				return
			case outCh <- order.NextActionState(actions):
			}
		}
	}()
	return outCh
}

func FinishStage(
	ctx context.Context,
	actions *model.OrderActions,
	orders <-chan model.OrderFinishedExternalInteraction,
) <-chan model.OrderProcessFinished {
	outCh := make(chan model.OrderProcessFinished, 10)
	go func() {
		defer close(outCh)
		for order := range orders {
			select {
			case <-ctx.Done():
				return
			case outCh <- order.NextActionState(actions):
			}
		}
	}()
	return outCh
}

func (o *OrderPipelineImplementation) Start(
	ctx context.Context,
	actions model.OrderActions,
	orders <-chan model.OrderInitialized,
	processed chan<- model.OrderProcessFinished,
) {
	startStageResult := StartStage(ctx, &actions, orders)
	processStageResult := ProcessStage(ctx, &actions, startStageResult)
	finishStageResult := FinishStage(ctx, &actions, processStageResult)

	for result := range finishStageResult {
		processed <- result
	}
}
