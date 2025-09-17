package main
import "fmt"
import "strings"

type person struct {
	name string
	age  int
}

func (p person) greet() string {
	return "Hello, my name is " + p.name
}

func (p *person) setName(newName string) {
	p.name = strings.ToUpper(newName)
}

func main() {
	var p1 = person{name: "wick", age: 21}
	fmt.Println(p1.greet())
	p1.setName("ethan")
	fmt.Println(p1.greet())
}