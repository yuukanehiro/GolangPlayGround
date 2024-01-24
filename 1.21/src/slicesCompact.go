package main

import (
	"fmt"
	"slices"
)

func main() {
	data := []int{1, 2, 2, 3, 3, 3, 4}
	compacted := slices.Compact(data) // 重複削除
	fmt.Println("Compacted:", compacted)
}

// Compacted: [1 2 3 4]
