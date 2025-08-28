package main
import ("fmt")

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	// type student struct {
    // 	name string
    // 	grade int
	// }

	// var s1 = student{}
	// s1.name = "wick"
	// s1.grade = 2

	// var s2 = student{"ethan", 2}

	// var s3 = student{name: "jason"}

	// fmt.Println("student 1 :", s1.name)
	// fmt.Println("student 2 :", s2.name)
	// fmt.Println("student 3 :", s3.name)

	var s1 = student{}
    s1.name = "wick"
    s1.age = 21
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("age   :", s1.age)
    fmt.Println("age   :", s1.person.age)
    fmt.Println("grade :", s1.grade)
}