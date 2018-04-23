package pointer

import (
	"errors"
	"fmt"
)

/*
	Comments may seem unnecessary here, but it's seems to be a
	style guideline. Otherwise, you'll have to deal with the
	"should have comment or be unexported" error.
*/

// Stringer : re-definition of a fmt function to configure a native format for a value. Remember toString() from Java? That's it.
type Stringer interface {
	String() string
}

// Bitcoin : a Bitcoin int alias
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet : structure defining a Bitcoin wallet and it's contents
type Wallet struct {
	balance Bitcoin
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// Deposit : add funds to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw : remove funds from a Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

// Balance : return the amount of funds a Wallet currently has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
