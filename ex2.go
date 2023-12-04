package main

import "fmt"

func MergeMaps(map1, map2 map[string]int) map[string]int {
	mergedMap := make(map[string]int)

	for key, value := range map1 {
		mergedMap[key] = value
	}

	for key, value := range map2 {
		mergedMap[key] = value
	}

	return mergedMap
}

func main() {
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 5, "c": 6, "d": 7}

	mergedMap := MergeMaps(map1, map2)

	fmt.Println("Mapa combinado:")
	for key, value := range mergedMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
