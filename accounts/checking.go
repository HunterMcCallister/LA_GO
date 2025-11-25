package accounts

import "LA_GO/customer"

// CheckingAccount represents a checking account that earns no interest.
type CheckingAccount struct {
	BaseAccount
}

// NewCheckingAccount creates a new checking account with number, customer, and balance.
func NewCheckingAccount(number string, c customer.Customer, balance float64) CheckingAccount {
	return CheckingAccount{
		BaseAccount: BaseAccount{
			Number:   number,
			Customer: c,
			Balance:  balance,
		},
	}
}

// Accrue for a checking account does nothing (just like the Java version).
func (ca *CheckingAccount) Accrue(rate float64) {
	// no interest
}

// AccrueAsync sends 0 interest for checking accounts.
func (ca *CheckingAccount) AccrueAsync(rate float64, ch chan float64) {
	ch <- 0
}
