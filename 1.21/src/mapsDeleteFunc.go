package main

import (
	"fmt"

	"maps"
)

func main() {
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	fmt.Printf("Original Map: %v\n", myMap)

	maps.DeleteFunc(myMap, func(k string, v int) bool {
		return k == "one"
	})

	fmt.Printf("with Deleted value %v\n", myMap)
}

// Original Map: map[four:4 one:1 three:3 two:2]
// with Deleted value map[four:4 three:3 two:2]
