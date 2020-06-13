package main

import (
	"errors"
	"fmt"
)

//Bitcoin is int
type Bitcoin int

//Wallet is
type Wallet struct {
	balance Bitcoin
}

//ErrInsufficientFunds is withdraw error
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Deposit is + coin
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance is coin
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw is - coin
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
