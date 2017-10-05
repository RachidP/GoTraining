package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("i=%v   %v \n", i, "FIZZBUZZ")
		case i%3 == 0:
			fmt.Printf("i=%v   %v \n", i, "FIZZ")
		case i%5 == 0:
			fmt.Printf("i=%v   %v \n", i, "BUZZ")
		default:
			fmt.Printf("i=%v   %v \n", i, i)
		}
	}
}
