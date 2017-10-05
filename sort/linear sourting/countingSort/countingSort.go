package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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

var path = "/Users/Rachid/workspace/Go/src/github.com/RachidP/GoTraining/go_design_pattern/sort/linear sourting/countingSort/countingSort.txt"
var heapSizeA int

func main() {
	k := 999
	//list := []int{-1, 2, 5, 3, 0, 2, 3, 0, 3}
	list := generateSlice(1000)
	listB := make([]int, len(list))
	countingSourt(list, listB, k)

	//fmt.Println(list)
	createFile()
	writeFile(list, "Lista prima dell'ordinamento:  ")
	//fmt.Printf("\n\n\nlista prima dell'ordinamento \n\n %v\n\n", list)
	writeFile(listB[1:], "Lista dopo l'ordinamento:  ")

	//fmt.Printf("lista Dopo dell'ordinamento \n %v\n", list[1:])

}
func countingSourt(listA, listB []int, k int) {
	listC := make([]int, k)
	/*
		fmt.Printf("listA = %v  size=%v \n", listA, len(listA))
		fmt.Printf("listB = %v  size=%v \n", listB, len(listB))
		fmt.Printf("listC = %v  size=%v \n", listC, len(listC))
	*/

	tmp := 0
	for j, sizelist := 1, len(listA); j < sizelist; j++ {
		tmp = listA[j]
		listC[tmp] = listC[tmp] + 1

	} // listC[i] now contains the number of elements equal to i .
	//	fmt.Printf("listC changed = %v  size=%v \n", listC, len(listC))

	for i := 1; i < k; i++ {

		listC[i] += listC[i-1]
	} //listC[i] now contains the number of elements less than or equal to i .

	//fmt.Printf("listC changed = %v  size=%v \n", listC, len(listC))

	for j := len(listA) - 1; j > 0; j-- {

		tempA := listA[j]
		tempC := listC[tempA]
		listB[tempC] = tempA
		listC[tempA]--
	}
	//	fmt.Printf("listB = %v  size=%v \n", listB, len(listB))
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

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		//	slice[i] = rand.Intn(999) - rand.Intn(999)
		slice[i] = rand.Intn(999-0) + 0
	}
	return slice
}
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
