package main

import (
	"fmt"
	"math"
)

func main() {

	number := float64(1024645454)
	fmt.Printf("numero=%v  radice quadrata = %v   \n-----------------------\n", number, math.Sqrt(number))
	fmt.Printf("\nrisultato %v   \n-----------------------\n", Sqrt(number))
	fmt.Printf("\nrisultato %v\n\n", SqrtDelta(number))

}

func Sqrt(x float64) float64 {

	tot := 1.0

	for n := 0; n < 1000; n++ {
		fmt.Printf("i=%v  , %v \n", n, tot)
		tot = tot - ((math.Pow(tot, 2) - x) / (2 * tot))

	}
	return tot
}

func SqrtDelta(x float64) float64 {

	tot := float64(1)
	delta := 0.0
	i := 0
	for delta != tot {
		delta = tot
		tot = tot - ((math.Pow(tot, 2) - x) / (2 * tot))
		fmt.Printf("i=%v, delta= %v  , %v \n", i, delta, tot)
		i++
	}
	return tot
}
