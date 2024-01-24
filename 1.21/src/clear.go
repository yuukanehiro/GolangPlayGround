package main

import "fmt"

func main() {
	myMap := map[string]int{"apple": 5, "banana": 3}
	clear(myMap)
	fmt.Println(myMap) // map[]
}
