package main

import (
	"fmt"

	"github.com/sony-nurdianto/GoRust/bank_programs/banks"
)

func main() {
	banks.Deposits(100)
	banks.Deposits(200)

	fmt.Println("Balance: ", banks.Balance())
}
