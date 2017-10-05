package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const infinito int = 1000000

var path = "/Users/Rachid/workspace/Go/src/github.com/RachidP/GoTraining/go_design_pattern/sort/heapSort/heapSort.txt"

func main() {

	//list := []int{5, 2, 4, 7, 1, 3, 86}
	list := generateSlice(3000)
	p := 1
	r := len(list)
	createFile()

	writeFile(list, "Lista prima dell'ordinamento:  ")
	mergeSort(list, p, r)
	// write the result into file

	writeFile(list, "Lista dopo l'ordinamento:  ")
	//	myMerge(list, p, q, r)

}

func myMerge(list []int, p int, q int, r int) {
	fmt.Println(list)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)

	n1 := q - p + 1
	n2 := r - q
	leftList := make([]int, n1+1)
	for i := 0; i < n1; i++ {
		leftList[i] = list[p+i-1]
	}
	leftList[n1] = infinito
	//fmt.Printf("l= %v \n", leftList)

	rightList := make([]int, n2+1)
	for j := 0; j < n2; j++ {
		rightList[j] = list[q+j]
	}
	rightList[n2] = infinito
	//fmt.Printf("l= %v \n", rightList)

	i, j := 0, 0
	for k := p - 1; k < r; k++ {
		if leftList[i] <= rightList[j] {
			list[k] = leftList[i]
			i++
		} else {
			list[k] = rightList[j]
			j++
		}

	}
	//fmt.Println(list)

}

func mergeSort(list []int, p int, r int) {

	if p < r {

		q := int(math.Floor(float64(p+r) / 2))
		mergeSort(list, p, q)
		mergeSort(list, q+1, r)
		myMerge(list, p, q, r)

	}
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
