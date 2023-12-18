package bank

import (
	"fmt"
	"strings"
	"time"
)

const (
	TopUpOp OperationType = iota
	WithdrawOp
)

type OperationType int64

type Clock interface {
	Now() time.Time
}

func NewRealTime() *RealClock {
	return &RealClock{}
}

type RealClock struct{}

func (c *RealClock) Now() time.Time {
	return time.Now()
}

type Operation struct {
	OpTime   time.Time
	OpType   OperationType
	OpAmount int
	Balance  int
}

func (o Operation) String() string {
	var format string
	if o.OpType == TopUpOp {
		format = `%s +%d %d`
	} else {
		format = `%s -%d %d`
	}
	return fmt.Sprintf(format, o.OpTime.String()[:19], o.OpAmount, o.Balance)
}

type Account interface {
	TopUp(amount int) bool
	Withdraw(amount int) bool
	Operations() []Operation
	Statement() string
	Balance() int
}

func NewAccount(clock Clock) Account {
	return &AccountImpl{
		operations: []Operation{},
		balance:    0,
		clock:      clock,
	}
}

type AccountImpl struct {
	operations []Operation
	balance    int
	clock      Clock
}

func (a *AccountImpl) addEvent(amount int, operation OperationType) {
	a.operations = append(a.operations, Operation{
		OpTime:   a.clock.Now(),
		OpType:   operation,
		OpAmount: amount,
		Balance:  a.balance,
	})
}

func (a *AccountImpl) TopUp(amount int) bool {
	if amount <= 0 || a.Balance() < 0 {
		return false
	}
	a.balance += amount
	a.addEvent(amount, TopUpOp)
	return true
}

func (a *AccountImpl) Withdraw(amount int) bool {
	if amount <= 0 || a.Balance() < amount {
		return false
	}
	a.balance -= amount
	a.addEvent(amount, WithdrawOp)
	return true
}

func (a *AccountImpl) Operations() []Operation {
	return a.operations
}

func (a *AccountImpl) Statement() string {
	result := make([]string, len(a.operations))
	for i, value := range a.operations {
		result[i] = value.String()
	}
	return strings.Join(result, "\n")
}

func (a *AccountImpl) Balance() int {
	return a.balance
}
