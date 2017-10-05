package main

import (
	"errors"
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

var path = "/Users/Rachid/workspace/Go/src/github.com/RachidP/GoTraining/go_design_pattern/sort/heapSort/heapSort.txt"
var heapSizeA int

func main() {

	list := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	//	list := generateSlice(100000000)
	//fmt.Println(list)
	//createFile()

	//writeFile(list, "Lista prima dell'ordinamento:  ")
	fmt.Printf("lista prima dell'ordinamento \n %v\n", list)
	heapSort(list)
	//writeFile(list, "Lista dopo l'ordinamento:  ")
	fmt.Printf("lista Dopo dell'ordinamento \n %v\n", list)

}

func maxHeapify(list []int, i int) {

	largest := i
	l := left(i)
	r := right(i)

	if l <= heapSizeA && list[l] > list[i] {

		largest = l

	}

	if r <= heapSizeA && list[r] > list[largest] {
		largest = r

	}
	if largest != i {

		list[i], list[largest] = list[largest], list[i]
		maxHeapify(list, largest)
	}
}

/*parent of child*/
func parent(i int) int {
	return (i - 1) / 2

}

/*left  node-child*/
func left(i int) int {

	return (2 * i) + 1
}

/*right node-child*/
func right(i int) int {

	return (2 * i) + 2
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

/*returns the element of the List with the largest key.*/
func heapMaximum(list []int) int {

	return list[0]
}

/*removes and returns the element of S with the largest key*/
func heapExtracMax(list []int) (int, error) {

	if heapSizeA < 0 {

		fmt.Println("heap underflow")
		return -1, errors.New("heap underflow")
	}

	max := list[0]
	list[0] = list[heapSizeA]
	heapSizeA--
	maxHeapify(list, 0)

	return max, nil
}

/*Increases the value of element x’s key to the new value k,
which is assumed to be at least as large as x’s current key value.
*/

func heapIncreaseKey(list []int, i int, key int) error {

	if key < list[i] {

		fmt.Println("New key is smaller than current key ")
		//return -1, errors.New("heap underflow")
		errObj := fmt.Errorf("My Error. New key is smaller than current key")
		return errObj

	}

	list[i] = key

	//oppure >1
	for i >= 0 && list[parent(i)] < list[i] {
		list[i], list[parent(i)] = list[parent(i)], list[i]

		i = parent(i)
	}

	fmt.Printf("\n Increases the value of element  x’s =%d key to the new value k=%d \n", key, i)
	//nessun errore
	return nil

}

/*Inserts the element x into the set S, which is equivalent to the operation S= S U {x} */

func maxHeapInsert(list []int, key int) error {

	heapSizeA++

	list = append(list, minInt)
	err := heapIncreaseKey(list, heapSizeA, key)
	if err != nil {
		fmt.Println("errore durante la chiamata heapIncreseKey dalla funzione maxHeapInsert")
		errObj := fmt.Errorf("errore durante la chiamata heapIncreseKey dalla funzione maxHeapInsert. Error: %s", err.Error())
		return errObj
	}

	return nil
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

/*********************************************/
func implementaPriorityQuee(list []int) {

	fmt.Printf("restituisce la chiave piu grande nella lsita \n %v\n", heapMaximum(list))

	fmt.Println("Elimina e restituisce l elemento nella lista con la chiave piu grande")
	if max, err := heapExtracMax(list); !isError(err) {
		fmt.Println("l'elemento eliminato e' ", max)
		fmt.Println("nuova lista :   ")
		fmt.Println(list)

	}

	/*
		tmp := 199
		fmt.Printf("aggiungi il valore %d alla lista \n", tmp)
		fmt.Println(list)

		if err := heapIncreaseKey(list, 0, tmp); !isError(err) {
			fmt.Println("l'elemento e stato aggiunto alla lista correttamente: tmp ", tmp)
			fmt.Println("nuova lista :   ")
			fmt.Println(list)

		}
	*/

}
func heapSort(list []int) {

	buildMaxHeap(list)

	dim := len(list) - 1

	for i := dim; i >= 1; i-- {

		list[0], list[i] = list[i], list[0]
		heapSizeA--

		maxHeapify(list, 0)

	}

}

func buildMaxHeap(list []int) {

	dim := len(list) - 1
	heapSizeA = dim
	//	tmp := int(math.Floor(float64(dim / 2)))
	tmp := dim / 2

	for i := tmp; i >= 0; i-- {
		maxHeapify(list, i)

	}

	//fmt.Println(list)
}
