package deposit

//Calculate рассчитывает параметры вклада.
func Calculate(amount int, currency string) (min int, max int) {
	minPercent, maxPercent := PercentFor(currency)
	min = ((amount * minPercent) / 100) / 12
	max = ((amount * maxPercent) / 100) / 12

	return min, max
}

// PercentFor возвращает процент по вкладу по валюте
func PercentFor(currency string) (min int, max int) {
	if currency == "TJS" {
		return 4, 6
	}

	if currency == "USD" {
		return 1, 2
	}

	return 0, 0 //TODO: что делать здесь?
}
