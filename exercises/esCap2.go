package main

import "fmt"

func main() {
	variAdic(7, 8, 9, 1, 2, 3, 4, 5)

	fmt.Printf("\n------------------MAP -----------------------------------------\n")
	m := []int{1, 2, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", mapF(f, m))

	fmt.Printf("\n------------------BUBBLE SORT -----------------------------------------\n")
	bubbleSort(9, 8, 7, 6, 99, 1, 33, 5, 65, 878, 45, 23, 355, 4)

	fmt.Printf("\n------------------retFunction-----------------------------------------\n")
	p2 := retFunction()
	fmt.Printf("return of function = %v\n", p2(2))
	fmt.Printf("return of function = %v\n", p2(3))
	p3 := plusX(3)
	fmt.Printf("return of function = %v\n", p3(22))
	fmt.Printf("return of function = %v\n", p3(33))
}

/*Write a function that takes a variable number of ints and prints each integer on a
separate line*/
func variAdic(numbers ...int) {

	for _, val := range numbers {

		fmt.Println(val)

	}

}

/*
Map function A map()-function is a function that takes a function and a
list. The function is applied to each member in the list and a new list containing these
calculated values is returned.*/

func mapF(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {

		j[k] = f(v)
	}
	return j

}

func bubbleSort(items ...int) {
	fmt.Printf("  items= %v \n", items)
	for {
		swapped := false
		for i, _ := range items {

			if i == 0 {
				i = 0
			} else {

				if items[i-1] > items[i] {
					items[i-1], items[i] = items[i], items[i-1]
					swapped = true
				}
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Printf("  items= %v \n", items)
}

/*Write a function that returns a function that performs a +2 on integers*/
func retFunction() func(int) int {

	f := func(v int) int {
		return v + 2
	}
	return f
}

/*Generalize the function from 1, and create a plusX(x) which returns functions
that add x to an integer*/
func plusX(x int) func(int) int {

	f := func(y int) int {
		return x + y
	}
	return f
}
