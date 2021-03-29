package card

import (
	"github.com/shodikhuja83/bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) (card types.Card) {
	card = types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active{
		if  amount < card.Balance{
			if amount <= 0 {
			const limit = 20_000_00
				if amount > limit {
					card.Balance -= amount
				}
			}
		}
	}
	return card
}
const Limit = 50_000_00
func Deposit(card *types.Card, amount types.Money){
if !card.Active || amount > Limit || amount < 0{
  return 
}
card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int){
	bonus := (int(card.MinBalance ) * percent / 100 * daysInMonth) / daysInYear	
	if !(*card).Active {
		return
	}
	if card.Balance < 0 {
		return
	}
	if bonus > 5_000_00 {
		return
	}
	card.Balance += types.Money(bonus)
}
//Total вычисляет общую сумму средств на всех картах.
//Отрицательные суммы и суммы на заблокированных картах игнорируются.
func Total(cards []types.Card) int {
	sum :=0
	for _, card := range cards {
		if (card.Balance<0 && !card.Active){
			return sum
		}
		sum += int(card.Balance)
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {

	var payment []types.PaymentSource
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance < 0 {
			continue
		}

		payment = append(payment, types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		})
	}
	return payment
}