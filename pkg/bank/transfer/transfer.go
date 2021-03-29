package transfer

func Total(amount int) (bonus int) {
	bonus = amount + (((amount * 5) / 10) / 100)
	return bonus
}
