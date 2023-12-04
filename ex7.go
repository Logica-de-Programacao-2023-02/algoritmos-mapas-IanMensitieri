package main

import "strings"

func WordLetterCountMap(frase string) map[string]map[rune]int {
	result := make(map[string]map[rune]int)
	palavras := strings.Fields(frase)

	for _, palavra := range palavras {
		letraCountMap := make(map[rune]int)
		for _, letra := range palavra {
			letraCountMap[letra]++
		}
		result[palavra] = letraCountMap
	}

	return result
}
