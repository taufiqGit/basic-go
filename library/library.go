package library

import "fmt"


func SayHello() {
	fmt.Println("Hello from the library!")
}

func introduce(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

type Student struct {
	Name  string
	Grade int
}