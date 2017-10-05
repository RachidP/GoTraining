package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//printA()
	str := "asSASA ddd dsjkdsjs dk"

	fmt.Printf("\n-----------------------------------------------------------\n")
	countStrin(str)

	fmt.Printf("\n-----------------------------------------------------------\n")
	changeStri(str)

	fmt.Printf("\n------------------Reverse-----------------------------------------\n")
	reverse(str)

	fmt.Printf("\n------------------AVERAGE  -----------------------------------------\n")
	slc := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	media(slc)

}

func media(slc []float64) {
	len := float64(len(slc))
	var tot, avg float64
	for val := range slc {
		tot += float64(val)
	}
	if tot == 0 {
		avg = 0

	} else {
		avg = tot / len
	}

	fmt.Printf("slc= %v sum=%v len=%v average=%v", slc, tot, len, avg)
}

/*count the element in the string*/
func countStrin(str string) {

	fmt.Printf("%s   len=%d  number of bytes in the string=%d \n", str, len(str), utf8.RuneCountInString(str))
}

/*Reverse of a string*/
func reverse(str string) {
	s := []byte(str)
	size := len(str)
	nS := make([]byte, size)

	var i int
	for size-1 > 0 {
		nS[i] = s[size-1]
		size--
		i++

	}
	fmt.Printf("string= %v  reverse string = %v", string(s), string(nS))

}

/*change the elements in the string*/
func changeStri(str string) {
	i := []rune{66, 1024, 65}
	r := string(i)
	fmt.Println(r)
	v := []byte(str)
	v[3] = 'a'
	v[4] = 'b'
	v[5] = 'c'
	v1 := string(v)
	fmt.Printf("old string =%s    new string= %v %v ", str, v, v1)
	copy(v[4:4+3], []byte("ZZZ"))
	fmt.Printf("old string =%s    new string= %v %v ", str, string(v), v1)
}

/* print a piramide of A*/
func printA() {

	temp := 0
	for i := 0; i < 100; i++ {
		for temp < i {
			fmt.Printf("A")
			temp++
		}
		fmt.Printf("\n")
		temp = 0
	}
	fmt.Printf("\n-----------------------------------------------------------\n")
	str := "A"
	for i := 0; i < 100; i++ {
		fmt.Printf("%s \n", str)
		str += "A"
	}

}
