package main

import "fmt"

func main() {
	var asd int = 23940
	stry := "itil mini"
	fmt.Println(asd, stry)

	var msg = ` hallo heheh "Taufiqs"
hahaha hdhe
jejrt jejrj`

	fmt.Println(msg)

	const (
		square = 939
		cube uint8  = 123
		isToday bool = true
	)

	fmt.Println(square, cube, isToday)

	var sx = "hello world"
	for i, v := range sx {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}

	var arr2 = [5]int{10, 20, 30, 40, 50}
	for _, v := range arr2 {
		fmt.Println("val", v)
	}

	var arr_str = [5]string{"a", "b", "c", "d", "e"}
	for i, v := range arr_str {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}
}