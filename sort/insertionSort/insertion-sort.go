package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var path = "/Users/Rachid/workspace/Go/src/github.com/RachidP/GoTraining/go_design_pattern/sort/test.txt"

func main() {

	list := generateSlice(20000)
	fmt.Println(list)
	insertionSort(list)

}
func insertionSort(list []int) {

	var key, i int
	dim := len(list)
	for j := 1; j < dim; j++ {
		key = list[j]

		i = j - 1
		for i >= 0 && list[i] >= key {
			list[i+1] = list[i]
			i--
		}
		list[i+1] = key
	}
	// write the result into file
	createFile()
	writeFile(list)

}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
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

func writeFile(list []int) {
	//var path = "/Users/novalagung/Documents/temp/test.txt"

	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)

	if isError(err) {
		return
	}
	defer file.Close()
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
