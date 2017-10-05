package main

import (
	"fmt"

	"time"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(2)

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i

		}

		//wg.Done()
	}()

	go func() {
		for {
			fmt.Println(<-c)

		}
		//wg.Done()

	}()

	//wg.Done()
	time.Sleep(time.Second)
}
