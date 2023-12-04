package main

import (
	"fmt"
	"strings"
)

func CountWordsOccurrences(text string) map[string]int {
	occurrences := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		occurrences[word]++
	}

	return occurrences
}

func main() {
	text := "Esta é uma string de exemplo. Esta string contém palavras repetidas, como exemplo."
	result := CountWordsOccurrences(text)

	fmt.Println("Ocorrências de palavras:")
	for word, count := range result {
		fmt.Printf("%s: %d\n", word, count)
	}
}
