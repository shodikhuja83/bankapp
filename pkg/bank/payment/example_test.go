package payment

import (
	"fmt"
	"github.com/shodikhuja83/bank/pkg/bank/types"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 0,
			Amount: 20_000,
		},
		{
			ID: 1,
			Amount: 20_000,
		},
		{
			ID: 2,
			Amount: 10_000,
		},
		{
			ID: 3,
			Amount: 30_000,
		},
	}

	max := Max(payments)
	fmt.Print(max.ID)
	// Output: 3
}