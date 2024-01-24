package main

import (
	"fmt"
	"slices"
)

func main() {
	original := []string{"a", "b", "c"}
	cloned := slices.Clone(original) // Sliceの浅いコピーを作成
	fmt.Println("Original:", original)
	fmt.Println("Cloned:", cloned)
}

// Original: [a b c]
// Cloned: [a b c]
