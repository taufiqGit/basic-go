package main
import ("fmt")

func main() {

	var numberPointer = 30
	fmt.Println("before :", numberPointer)

	change(&numberPointer, 100)
	fmt.Println("after :", numberPointer)

	var ungker = 30
	var ihi *int = &ungker
	fmt.Println(ungker)
	fmt.Println(ihi)
}

func change(original *int, value int) {
	*original = value
}
