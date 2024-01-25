package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var newSingletonOnce = sync.OnceValue(func() *singleton {
	return &singleton{}
})

func newSingleton() *singleton {
	return newSingletonOnce()
}

func main() {
	instance1 := newSingleton()
	instance2 := newSingleton()
	fmt.Println(instance1 == instance2)
}

// true
// 説明: 一度だけインスタンスを生成する
