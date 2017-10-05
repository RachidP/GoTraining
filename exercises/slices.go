package main

import (
	"fmt"
)

func main() {

	slc1 := make([]int, 8, 10)
	fmt.Println(slc1, len(slc1), cap(slc1))
	slc1 = slc1[:cap(slc1)]
	fmt.Println(slc1, len(slc1), cap(slc1))

}
