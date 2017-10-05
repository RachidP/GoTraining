package main

import (
	"fmt"
	"os"
)

const (
	minUint uint = 0 // binary: all zeroes

	// Perform a bitwise NOT to change every bit from 0 to 1
	maxUint = ^minUint // binary: all ones

	// Shift the binary number to the right (i.e. divide by two)
	// to change the high bit to 0
	maxInt = int(maxUint >> 1) // binary: all ones except high bit

	// Perform another bitwise NOT to change the high bit to 1 and
	// all other bits to 0
	minInt = ^maxInt // binary: all zeroes except high bit
)

var path = "/Users/Rachid/workspace/Go/src/github.com/RachidP/GoTraining/go_design_pattern/sort/heapSort/heapSort.txt"
var heapSizeA int

func main() {

	list := []int{-1, 2, 8, 7, 1, 3, 5, 6, 4}
	//	list := generateSlice(100000000)
	//fmt.Println(list)
	//createFile()

	//writeFile(list, "Lista prima dell'ordinamento:  ")
	fmt.Printf("lista prima dell'ordinamento \n %v\n", list)
	quickSort(list, 1, len(list)-1)
	//partition(list, 1, len(list)-1)
	//writeFile(list, "Lista dopo l'ordinamento:  ")
	fmt.Printf("lista Dopo dell'ordinamento \n %v\n", list[1:])

}

func quickSort(list []int, p int, r int) {

	if p < r {

		q := partition(list, p, r)
		quickSort(list, p, q-1)
		quickSort(list, q+1, r)

	}
}

func partition(list []int, p int, r int) int {

	x := list[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		fmt.Println(list)
		if list[j] <= x {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[r] = list[r], list[i+1]
	return i + 1

}

/****************************************WRITING AND READING INTO FILE ********************************************/
func createFile() {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
}

func writeFile(list []int, desc string) {
	//var path = "/Users/novalagung/Documents/temp/test.txt"

	// open file using READ & WRITE permission 0644
	//var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)

	if isError(err) {
		return
	}
	defer file.Close()
	// write into file
	_, err = file.WriteString(fmt.Sprintln(desc))
	if isError(err) {
		return
	}

	// write into file
	_, err = file.WriteString(fmt.Sprintln(list))
	if isError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> done writing to file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
