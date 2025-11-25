package accounts

import "LA_GO/customer"

// SavingAccount represents a savings account that earns interest.
type SavingAccount struct {
	BaseAccount
	Interest float64
}

// NewSavingAccount creates a new savings account with number, customer, and balance.
func NewSavingAccount(number string, c customer.Customer, balance float64) SavingAccount {
	return SavingAccount{
		BaseAccount: BaseAccount{
			Number:   number,
			Customer: c,
			Balance:  balance,
		},
		Interest: 0,
	}
}

// Accrue adds interest to both the interest field and the balance.
// (This is the simple, Phase 1 version.)
func (sa *SavingAccount) Accrue(rate float64) {
	earned := sa.Balance * rate
	sa.Interest += earned
	sa.Balance += earned
}

// AccrueAsync runs interest calculation in a goroutine by sending the earned interest over a channel.
// (This is the Phase 2 version.)
func (sa *SavingAccount) AccrueAsync(rate float64, ch chan float64) {
	earned := sa.Balance * rate
	sa.Interest += earned
	sa.Balance += earned

	// Send the earned interest back to the bank
	ch <- earned
}
