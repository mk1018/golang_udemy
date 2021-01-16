package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)

	ch1 <- 1

	close(ch1)

	i, ok := <-ch1
	fmt.Println(i, ok)
}
