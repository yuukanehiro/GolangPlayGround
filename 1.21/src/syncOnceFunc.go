package main

import (
	"fmt"
	"sync"
)

var onceFunc = sync.OnceFunc(func() {
	fmt.Println("Hello")
})

func main() {
	onceFunc()
	onceFunc()
}

// Hello
// 説明: 一度だけ関数がされた
