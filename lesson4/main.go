package main

import "fmt"

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)

	fmt.Printf("%T, %T\n", fl64, fl)

	fmt.Printf("%T\n", float32(fl64))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

}
