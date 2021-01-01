package wallet

import (
	"errors"
	"fmt"
)

//Bitcoin is a cryptocurrency
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet holds cash amounts
type Wallet struct {
	balance Bitcoin
}

// Deposit accepts an amount and credits the balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds is an error message
var ErrInsufficientFunds = errors.New("insufficient funds")

// Withdraw accepts an amount and debits the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
