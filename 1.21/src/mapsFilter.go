package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	sampleMap := map[string]int{"apple": 5, "banana": 3, "cherry": 7}
	filteredMap := maps.Filter(sampleMap, func(k string, v int) bool {
		return v > 4
	})
	fmt.Println("Filtered map:", filteredMap)
}
