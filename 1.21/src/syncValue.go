package main

import (
	"fmt"
	"sync"
)

var count int

// sync.OnceValueはsync.Onceと同様に、一度だけ関数を実行する。
var onceValue = sync.OnceValue(func() int {
	count++
	return count
})

func main() {
	fmt.Println(onceValue())
	fmt.Println(onceValue())
}

// 1
// 1
