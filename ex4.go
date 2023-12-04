package main

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func GroupAnagrams(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		sortedWord := SortString(word)
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	return anagramGroups
}
