package card

import (
	"github.com/shodikhuja83/bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}
func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 0)
	fmt.Println(result.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}
func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 0000",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 1111",
			Balance: 15_000_00,
			Active:  false,
		},
	}
	paymentSources := PaymentSources(cards)

	for _, v := range paymentSources {
		fmt.Println(v.Number)
	}

	//Output:  5058 xxxx xxxx 8888	
}
func ExampleDeposit_positiv() {
	cardtaking := types.Card{Balance: 10_000_00, Active: true}
	Deposit(&cardtaking, 50_000_00)
	fmt.Println(cardtaking.Balance)
	//Output: 6000000
}
func ExampleDeposit_negotiveBalance() {
	cardtaking := types.Card{Balance: - 10_000_00, Active: true}
	Deposit(&cardtaking, 50_000_00)
	fmt.Println(cardtaking.Balance)
	//Output: 4000000
}

func ExampleDeposit_negative() {
	cardtaking := types.Card{Balance: - 10_000_00, Active: true}
	Deposit(&cardtaking, 20_000_00)
	fmt.Println(cardtaking.Balance)
	//Output: 1000000
}

func ExampleDeposit_inactive() {
	cardtaking := types.Card{Balance: 10_000_00, Active: false}
	Deposit(&cardtaking, 50_000_00)
	fmt.Println(cardtaking.Balance)
	//Output: 1000000
}

func ExampleDeposit_limitEqual() {
  cardtaking := types.Card{Balance: 40_000_00, Active: true}
	Deposit(&cardtaking, 40_000_00)
	fmt.Println(cardtaking.Balance)
	//Output: 8000000
}

func ExampleDeposit_limitAbove() {
	cardtaking := types.Card{Balance: 10_000_00, Active: true}
	Deposit(&cardtaking, 50_000_00)
	fmt.Println(cardtaking.Balance)
	//Output: 6000000
}
func ExampleDeposit_limit() {
	cardtaking := types.Card{Balance: 10_000_00, Active: true}
	Deposit(&cardtaking, 50_000_00)
	fmt.Println(cardtaking.Balance)
	//Output: 6000000
}

func ExampleAddBonus_positive() {
    bonus := types.Card{Balance: 10_000_00, Active: true}
	AddBonus(&bonus, 3, 30, 365)
	fmt.Println(bonus.Balance)
    // Output: 1000000
}

func ExampleAddBonus_inactive() {
    bonus := types.Card{Balance: 10_000_00, Active: false}
	AddBonus(&bonus, 3, 30, 365)
	fmt.Println(bonus.Balance)
    // Output: 1000000
}

func ExampleAddBonus_noMoney() {
    bonus := types.Card{Balance: 10_000_00, Active: true}
	AddBonus(&bonus, 3, 30, 365)
	fmt.Println(bonus.Balance)
    // Output: 1000000
}

func ExampleAddBonus_limit() {
    bonus := types.Card{Balance: 10_000_00, Active: true}
	AddBonus(&bonus, 3, 30, 365)
	fmt.Println(bonus.Balance)
    // Output: 1000000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	}
	fmt.Println(Total(cards))
	// Output: 3000000
}