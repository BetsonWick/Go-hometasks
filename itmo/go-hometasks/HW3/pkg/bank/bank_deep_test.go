package bank

import (
	"math"
	"testing"
)

func Test_balance_Upped(t *testing.T) {
	var account = NewAccount(NewMockTime())
	account.TopUp(1000)
	TestCase_balance{
		account:     account,
		wantBalance: 1000,
	}.Run(t)
}

func Test_topUp_Account_Large(t *testing.T) {
	acc := NewAccount(NewMockTime())
	acc.TopUp(math.MaxInt16)
	TestCase_topUp{
		account:     acc,
		amount:      math.MaxInt16,
		wantIsTopUp: true,
		wantBalance: math.MaxInt16 << 1,
	}.Run(t)
}

func Test_withdraw_Account_Large(t *testing.T) {
	acc := NewAccount(NewMockTime())
	acc.TopUp(math.MaxInt32)
	TestCase_withdraw{
		account:        acc,
		amount:         math.MaxInt16,
		wantIsWithdraw: true,
		wantBalance:    math.MaxInt32 - math.MaxInt16,
	}.Run(t)
}

func Test_operations_ManyOperations_Large(t *testing.T) {
	clock := NewMockTime()
	TestCase_operations{
		account: NewAccount(NewMockTime()),
		operations: []TestOperations{
			{
				opType: TopUpOp,
				amount: math.MaxInt16,
			},
			{
				opType: WithdrawOp,
				amount: math.MaxInt32,
			},
			{
				opType: WithdrawOp,
				amount: math.MaxInt16,
			},
		},
		wantOperations: []Operation{
			{
				OpTime:   clock.Now(),
				OpType:   TopUpOp,
				OpAmount: 32767,
				Balance:  32767,
			},
			{
				OpTime:   clock.Now(),
				OpType:   WithdrawOp,
				OpAmount: 32767,
				Balance:  0,
			},
		},
	}.Run(t)
}

func Test_statement_ManyOperations_Large(t *testing.T) {
	TestCase_statement{
		account: NewAccount(NewMockTime()),
		operations: []TestOperations{
			{
				opType: TopUpOp,
				amount: math.MaxInt16,
			},
			{
				opType: WithdrawOp,
				amount: math.MaxInt32,
			},
			{
				opType: WithdrawOp,
				amount: math.MaxInt16,
			},
		},
		wantStatement: `2023-03-18 12:34:07 +32767 32767
2023-03-18 12:34:37 -32767 0`,
	}.Run(t)
}
