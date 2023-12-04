package main

func SumMapValues(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}
