package main

import "dasar-golang/library"

var mhs1 = library.Student{
	Name:  "Taufiqs",
	Grade: 90,
}
func main() {
	library.SayHello()
	fmt.Println(mhs1)
	//library.introduce("Taufiqs", 21)
}