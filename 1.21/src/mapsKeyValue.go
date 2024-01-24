package main

import (
	"fmt"

	"golang.org/x/exp/maps" // mapsパッケージは標準だけど、KeysとValuesはexpパッケージ
)

func main() {
	sampleMap := map[string]int{"apple": 5, "banana": 3, "cherry": 7}
	keys := maps.Keys(sampleMap)
	values := maps.Values(sampleMap)
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)
}

// Keys: [apple banana cherry]
// Values: [5 3 7]
