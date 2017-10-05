package main

import "fmt"

const (
	val1 = iota

	val2
	val3
	val4
	val5
)

func main() {

	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(val3)
	fmt.Println(val4)
	fmt.Println(val5)
}
