package main

import "fmt"

func main() {
	a := 1

	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	}

	if b := 100; b == 100 {
		fmt.Println("a")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
}
