package main
import ("fmt")

func main() {

	var numberPointer = 30
	fmt.Println("before :", numberPointer)

	change(&numberPointer, 100)
	fmt.Println("after :", numberPointer)
}

func change(original *int, value int) {
	*original = value
}