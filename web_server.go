package main

import "fmt"
import "net/http"
type person struct {
	name string
	age  int
}

func index(w http.ResponseWriter, r *http.Request) {
	var p1 = person{name: "wick", age: 21}
	fmt.Fprintf(w, "name : %s\n", p1.name)
	fmt.Fprintf(w, "age  : %d\n", p1.age)
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "halo!")
    })

    http.HandleFunc("/index", index)

    fmt.Println("starting web server at http://localhost:3002/")
    http.ListenAndServe(":3002", nil)
}