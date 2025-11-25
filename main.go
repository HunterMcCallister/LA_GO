package main

import (
	"LA_GO/accounts"
	"LA_GO/bank"
	"LA_GO/customer"
	"fmt"
)

func main() {
	b := bank.NewBank()

	c := customer.NewCustomer("Ann")

	b.Add(&accounts.CheckingAccount{
		BaseAccount: accounts.BaseAccount{
			Number:   "01001",
			Customer: c,
			Balance:  100.00,
		},
	})

	b.Add(&accounts.SavingAccount{
		BaseAccount: accounts.BaseAccount{
			Number:   "01002",
			Customer: c,
			Balance:  200.00,
		},
		Interest: 0,
	})

	b.Accrue(0.02)

	fmt.Println(b)
}
