package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println(rand.Intn(100))
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("random 1", randomizer.Int())
	fmt.Println("random 2", randomizer.Int())
	fmt.Println("random 3", randomizer.Float64())
	fmt.Println("random string", randomString(10))
}
