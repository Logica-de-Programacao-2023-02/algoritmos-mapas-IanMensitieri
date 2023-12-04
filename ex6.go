package main

func SumWordCounts(listOfMaps []map[string]int) map[string]int {
	sumMap := make(map[string]int)

	for _, wordCountMap := range listOfMaps {
		for word, count := range wordCountMap {
			sumMap[word] += count
		}
	}

	return sumMap
}
