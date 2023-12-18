package model

type OrderState string

const (
	Initialized                 OrderState = "order_initialized"
	ProcessStarted              OrderState = "order_process_started"
	FinishedExternalInteraction OrderState = "order_finished_external_interaction"
	ProcessFinished             OrderState = "order_process_finished"
)

type OrderActions struct {
	InitToStarted                                func()
	StartedToFinishedExternalInteraction         func()
	FinishedExternalInteractionToProcessFinished func()
}

type OrderStates interface {
	OrderInitialized | OrderProcessStarted | OrderFinishedExternalInteraction | OrderProcessFinished
}

type OrderInitialized struct {
	OrderID     int
	ProductID   int
	OrderStates []OrderState
	Error       error
}

type OrderProcessStarted struct {
	OrderInitialized OrderInitialized
	OrderStates      []OrderState
	Error            error
}

type OrderFinishedExternalInteraction struct {
	OrderProcessStarted OrderProcessStarted
	StorageID           int
	PickupPointID       int
	OrderStates         []OrderState
	Error               error
}

type OrderProcessFinished struct {
	OrderFinishedExternalInteraction OrderFinishedExternalInteraction
	OrderStates                      []OrderState
	Error                            error
}

type Order struct {
	OrderID       int
	ProductID     int
	StorageID     int
	PickupPointID int
	IsProcessed   bool
	OrderStates   []OrderState
}

type OrderActionState[T OrderStates] interface {
	OrderInitialized | OrderProcessStarted | OrderFinishedExternalInteraction
	NextActionState(actions *OrderActions) T
}

func (ord OrderInitialized) NextActionState(actions *OrderActions) OrderProcessStarted {
	if ord.Error == nil {
		actions.InitToStarted()
		ord.OrderStates = append(ord.OrderStates, ProcessStarted)
	}
	return OrderProcessStarted{
		ord,
		ord.OrderStates,
		ord.Error,
	}
}

func (ord OrderProcessStarted) NextActionState(actions *OrderActions) OrderFinishedExternalInteraction {
	if ord.Error == nil {
		actions.StartedToFinishedExternalInteraction()
		ord.OrderStates = append(ord.OrderStates, FinishedExternalInteraction)
	}
	return OrderFinishedExternalInteraction{
		ord,
		ord.OrderInitialized.ProductID%2 + 1,
		ord.OrderInitialized.ProductID%3 + 1,
		ord.OrderStates,
		ord.Error,
	}
}

func (ord OrderFinishedExternalInteraction) NextActionState(actions *OrderActions) OrderProcessFinished {
	if ord.Error == nil {
		actions.FinishedExternalInteractionToProcessFinished()
		ord.OrderStates = append(ord.OrderStates, ProcessFinished)
	}
	return OrderProcessFinished{
		ord,
		ord.OrderStates,
		ord.Error,
	}
}

func CreateEmptyOrder(num int) OrderInitialized {
	return OrderInitialized{
		num,
		num,
		make([]OrderState, 0),
		nil,
	}
}

func CreateEmptyActions() OrderActions {
	return OrderActions{
		func() {},
		func() {},
		func() {},
	}
}
