package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i

		}
		close(c)
	}()

	go func() {
		for val := range c {
			fmt.Println(val)
		}
	}()
	time.Sleep(time.Second)
}
