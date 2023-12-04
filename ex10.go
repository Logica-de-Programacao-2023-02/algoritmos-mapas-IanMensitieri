package main

import "fmt"

func CountPairsFrequency(numbers []int) map[string]int {
	pairFrequency := make(map[string]int)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := fmt.Sprintf("%d:%d", numbers[i], numbers[j])
			pairFrequency[pair]++
		}
	}

	return pairFrequency
}
