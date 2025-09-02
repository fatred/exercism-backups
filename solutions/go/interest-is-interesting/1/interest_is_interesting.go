package interest

// 3.213% for a balance less than 0 dollars (balance gets more negative).
const belowZero = 3.213

// 0.5% for a balance greater than or equal to 0 dollars, and less than 1000 dollars.
const zeroTo1k = 0.5

// 1.621% for a balance greater than or equal to 1000 dollars, and less than 5000 dollars.
const oneKto5K = 1.621

// 2.475% for a balance greater than or equal to 5000 dollars.
const fiveKorMore = 2.475

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    if balance < 0.0 {
		return belowZero
	} else if (balance >= 0.0) && (balance < 1000.0) {
		return zeroTo1k
	} else if (balance >= 1000.0) && (balance < 5000.0) {
		return oneKto5K
	} else if balance >= 5000.0 {
		return fiveKorMore
	} else {
		return 0.0
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + (Interest(balance))
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
    for years = 0; balance < targetBalance; years++ {
        balance = balance + (Interest(balance))
    }
	return years
}
