# Language Assignment (Go)

This project is a Go port of the simple banking application originally written in Java.  
It follows the assignment requirements by using multiple Go packages and implementing
a second version of interest accrual using goroutines and a channel.


## Overview

**Packages:**
- `customer` – defines the Customer type.
- `accounts` – defines the Account interface and the account types.
- `bank` – stores accounts and applies interest.
- `main` – builds a sample bank and runs the program.

The bank stores accounts in a `map[Account]struct{}`, similar to a Java `HashSet`.  
Interest calculation is run through goroutines, and each account sends its earned interest over a channel.  
The bank collects these values and prints the total interest accrued.

## Running the Program

From the command line run
`go run .`


