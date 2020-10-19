package main

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is generic error for insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is count of bitcoins
type Bitcoin float64

// Stringer method for Bitcoin
type Stringer interface {
	String() string
}

// Wallet defines a bitcoin wallet
type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

// Balance returns the balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit adds to the balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw removes from the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func main() {

}
