package main

func EqualizeExpenses(expenses map[string]float64) map[string]float64 {
	totalExpense := 0.0

	for _, expense := range expenses {
		totalExpense += expense
	}

	averageExpense := totalExpense / float64(len(expenses))

	equalizedExpenses := make(map[string]float64)
	for person, expense := range expenses {
		equalizedExpenses[person] = expense - averageExpense
	}

	return equalizedExpenses
}
