package utils


func CurrencyConversion(origin string, destiny string, amount float64) float64 {
	var fee float64 = 2/100
	var iof float64 =  1.1/100
	var totalAmount = amount - amount * fee
	totalAmount = totalAmount - totalAmount * iof

	if origin == "BRL" && destiny == "USD" {
		totalAmount = totalAmount/5.49
		return totalAmount
	}

	if origin == "BRL" && destiny == "EUR" {
		totalAmount = totalAmount/6.31
		return totalAmount
	}

	if origin == "BRL" && destiny == "GBP" {
		totalAmount = totalAmount/7.37
		return totalAmount
	}

	if origin == "USD" && destiny == "BRL" {
		totalAmount = totalAmount * 5.49
		return totalAmount
	}

	if origin == "EUR" && destiny == "BRL" {
		totalAmount = totalAmount *  6.31
		return totalAmount
	}

	if origin == "GBP" && destiny == "BRL" {
		totalAmount = totalAmount * 7.37
		return totalAmount
	}
	return amount
}