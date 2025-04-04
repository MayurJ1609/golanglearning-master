package closures

func ActivateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}
