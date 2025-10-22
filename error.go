package main

import (
	"errors"
	"fmt"
	"strings"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input cannot be empty")
	}
	return true, nil
}

func main() {
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	// var number int
	// var err error
	// number, err = strconv.Atoi(input)

	if valid, err := validate(input); valid {
		fmt.Println("hall")
	} else {
		panic(err.Error())
		fmt.Println("enddd")
	}
	// if err == nil {
	//     fmt.Println(number, "is number")
	// } else {
	//     fmt.Println(input, "is not number")
	//     fmt.Println(err.Error())
	// }
}
