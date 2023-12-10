package main

import "fmt"

func main() {
	var toplam int = 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			toplam += i
		} else if i%5 == 0 {
			toplam += i
		}
	}

	fmt.Println(toplam)
}
