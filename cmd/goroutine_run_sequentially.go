package main

import "fmt"

func main() {
	c1Begin := make(chan int, 1)
	// c1End := make(chan int, 1)
	c2Begin := make(chan int, 1)
	// c2End := make(chan int, 1)
	c3Begin := make(chan int, 1)
	c3End := make(chan int, 1)

	go func() {
		for {
			cmd := <-c1Begin
			if cmd == 0 {
				c2Begin <- 0
				break
			}

			fmt.Printf("cat\n")

			c2Begin <- 1
		}
	}()

	go func() {
		for {
			cmd := <-c2Begin
			if cmd == 0 {
				c3Begin <- 0
				break
			}

			fmt.Printf("dog\n")

			c3Begin <- 1
		}
	}()

	go func() {
		for {
			cmd := <-c3Begin
			if cmd == 0 {
				c3Begin <- 0
				break
			}
			fmt.Printf("fish\n")

			c3End <- 1
		}

	}()

	for i := 0; i < 5; i++ {
		c1Begin <- 1

		<-c3End
	}

	c1Begin <- 0
}
