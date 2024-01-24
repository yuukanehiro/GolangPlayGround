package main

import (
	"fmt"
	"slices"
)

func main() {
	data := []int{1, 3, 5, 7, 9}
	index, found := slices.BinarySearch(data, 5)
	fmt.Printf("Found 5 at index %d: %v\n", index, found) // Found 5 at index 2: true
}
