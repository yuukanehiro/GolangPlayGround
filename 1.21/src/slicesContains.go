package main

import (
	"fmt"
	"slices"
)

func main() {
	data := []string{"apple", "banana", "cherry"}
	contains := slices.Contains(data, "banana") // Sliceに指定した要素が含まれているかどうか真偽値を返却
	fmt.Println("Contains banana:", contains)
}

// Contains banana: true
