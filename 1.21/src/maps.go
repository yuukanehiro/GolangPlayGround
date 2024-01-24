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

	myOtherMap := maps.Clone(myMap)

	fmt.Printf("First Map %v \nClone %v\n", myMap, myOtherMap)
	// First Map map[four:4 one:1 three:3 two:2]
	// Clone map[four:4 one:1 three:3 two:2]

	resutl := maps.Equal(myMap, myOtherMap)
	fmt.Println("Equal?", resutl) // Equal? true
}
