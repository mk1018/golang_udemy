package main

import "fmt"

func main() {
	// 全ての型を表すだけで演算はできない
	var x interface{}

	x = 2
	fmt.Println(x)

	fmt.Println(x)
}
