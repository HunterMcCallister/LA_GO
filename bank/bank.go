package bank

import (
	"LA_GO/accounts"
	"fmt"
)

// Bank holds a set of accounts.
// We use map[Account]struct{} to mimic a HashSet-like structure.
type Bank struct {
	Accounts map[accounts.Account]struct{}
}

// NewBank creates a new Bank with an initialized map.
func NewBank() Bank {
	return Bank{
		Accounts: make(map[accounts.Account]struct{}),
	}
}

// Add inserts an account into the bank.
func (b *Bank) Add(a accounts.Account) {
	b.Accounts[a] = struct{}{}
}

// Accrue applies interest to every account using goroutines and a channel.
// It collects all earned interest and prints the total.
func (b *Bank) Accrue(rate float64) {
	ch := make(chan float64)

	// Launch one goroutine per account
	for acct := range b.Accounts {
		go acct.AccrueAsync(rate, ch)
	}

	var total float64
	count := len(b.Accounts)

	// Collect results
	for i := 0; i < count; i++ {
		total += <-ch
	}

	fmt.Printf("Total interest accrued: %.2f\n", total)
}

// String returns all accounts, each on its own line.
func (b Bank) String() string {
	result := ""
	for acct := range b.Accounts {
		result += fmt.Sprintf("%s\n", acct.String())
	}
	return result
}
