package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i

		}
		done <- true
	}()
	go func() {
		for i := 10; i < 20; i++ {
			c <- i

		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(done)
		close(c)
	}()

	go func() {
		for val := range c {
			fmt.Println(val)
		}
	}()
	time.Sleep(time.Second)
}
