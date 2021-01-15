package main

import "fmt"

func main() {
	var i int = 100
	// int8 , int16 , int32 , int64が用意されている

	var i2 int64 = 200

	fmt.Println(i + 50)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))
}
