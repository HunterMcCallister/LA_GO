package accounts

import (
	"LA_GO/customer"
	"fmt"
)

// Account is the interface that all account types must implement.
// It matches the abstract methods in the Java Account class.
type Account interface {
	Accrue(rate float64)
	AccrueAsync(rate float64, ch chan float64)
	fmt.Stringer
}

// BaseAccount holds shared fields for all account types.
// CheckingAccount and SavingAccount will embed this.
type BaseAccount struct {
	Number   string
	Customer customer.Customer
	Balance  float64
}

// Deposit adds money to the account balance.
func (ba *BaseAccount) Deposit(amount float64) {
	ba.Balance += amount
}

// Withdraw subtracts money from the account balance.
func (ba *BaseAccount) Withdraw(amount float64) {
	ba.Balance -= amount
}

// BalanceValue returns the current balance.
// (Named differently so it doesn't conflict with the Balance field.)
func (ba *BaseAccount) BalanceValue() float64 {
	return ba.Balance
}

// String returns the account as a formatted string.
func (ba BaseAccount) String() string {
	return fmt.Sprintf("%s:%s:%0.2f", ba.Number, ba.Customer, ba.Balance)
}
