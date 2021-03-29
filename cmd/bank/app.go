package main

import (
	"github.com/shodikhuja83/bank/pkg/bank/deposit"
	"github.com/shodikhuja83/bank/pkg/bank/transfer"
	"fmt"
)

func execute() {
	fmt.Println(deposit.Calculate(0, "TJS"))
	fmt.Println(deposit.Calculate(0, "USD"))
	fmt.Println(deposit.Calculate(500_000_00, "TJS"))
	fmt.Println(deposit.Calculate(500_000_00, "USD"))
	fmt.Println(deposit.Calculate(1_000_000_00, "TJS"))
	fmt.Println(deposit.Calculate(1_000_000_00, "USD"))
	fmt.Println(transfer.Total(0))
	fmt.Println(transfer.Total(5_000_00))
	fmt.Println(transfer.Total(10_000_00))
}