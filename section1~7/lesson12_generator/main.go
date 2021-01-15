package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	its := integers()
	fmt.Println(its())
	fmt.Println(its())
	fmt.Println(its())

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
}
