package main

import (
	"fmt" 
	"math/rand" 
	"time"
	"math"
)

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomWithRange(min, max int) int {
	return randomizer.Intn(max-min) + min
}

func calculate(d float64) (float64, float64) {
    // hitung luas
    var area = math.Pi * math.Pow(d / 2, 2)
    // hitung keliling
    var circumference = math.Pi * d

    // kembalikan 2 nilai
    return area, circumference
}

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

	var buah = []string{"nanas", "pear", "mango", "klengkeng", "pisang"}
	fmt.Println("panjang array/slice buah:", len(buah))
	fmt.Println("capasitas array/slice buah:", cap(buah))
	
	dst := make([]string, 5)
	src := []string{"apel", "jeruk", "anggur", "kiwi", "semangka"}
	n := copy(dst, src)
	fmt.Println("dst:", dst)
	fmt.Println("src:", src)
	fmt.Println("jumlah elemen yang disalin:", n)

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["ayam kampung"] = 10
	chicken["ayam potong"] = 20
	fmt.Println("jumlah ayam kampung:", chicken["ayam kampung"])
	fmt.Println("jumlah ayam potong:", chicken["ayam potong"])
	fmt.Println("nor", chicken["soto ayam"])

	randomValue := randomWithRange(1, 100)
	fmt.Println("random value:", randomValue)

	var diameter float64 = 10.0
	area, circumference := calculate(diameter)
	fmt.Printf("Luas lingkaran dengan diameter %.2f adalah %.2f\n", diameter, area)
	fmt.Printf("Keliling lingkaran dengan diameter %.2f adalah %.2f\n", diameter, circumference) 
}