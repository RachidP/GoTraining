package main

import (
	"fmt"
	"os"
)

var path = "/Users/Rachid/workspace/Go/src/github.com/RachidP/GoTraining/go_design_pattern/sort/heapSort/heapSort.txt"
var heapSizeA int

func main() {
	var top int
	list := make([]int, 10)
	fmt.Println(list)
	push(list, &top, 5)
	push(list, &top, 15)
	push(list, &top, 25)
	push(list, &top, 35)
	push(list, &top, 15)
	push(list, &top, 25)
	push(list, &top, 35)
	push(list, &top, 15)
	push(list, &top, 25)
	push(list, &top, 35)
	fmt.Printf("top=%d\nlista= %v\n", top, list)

	if val, err := pop(list, &top); !isError(err) {
		fmt.Println("l'elemento eliminato e' ", val)
		fmt.Printf("top=%d\n lista= %v\n", top, list)
	}

	if val, err := pop(list, &top); !isError(err) {
		fmt.Println("l'elemento eliminato e' ", val)
		fmt.Printf("top=%d\n lista= %v\n", top, list)
	}

	if val, err := pop(list, &top); !isError(err) {
		fmt.Println("l'elemento eliminato e' ", val)
		fmt.Printf("top=%d\n lista= %v\n", top, list)
	}

	if val, err := pop(list, &top); !isError(err) {
		fmt.Println("l'elemento eliminato e' ", val)
		fmt.Printf("top=%d\n lista= %v\n", top, list)
	}

	if val, err := pop(list, &top); !isError(err) {
		fmt.Println("l'elemento eliminato e' ", val)
		fmt.Printf("top=%d\n lista= %v\n", top, list)
	}

}

func isEmpty(list []int, top int) bool {
	if top == 0 {
		return true
	}
	return false

}

func push(list []int, top *int, x int) {
	(*top)++
	//list[*top] = x
	list = append(list, x)

}

func pop(list []int, top *int) (int, error) {
	if isEmpty(list, *top) {

		fmt.Println("Stack underflow. the stack is empty.")
		errObj := fmt.Errorf("Stack underflow. the stack is empty")
		return -1, errObj
	}
	(*top)--
	return list[(*top)+1], nil
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
