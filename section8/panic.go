package main

import "fmt"

// プログラムを強制終了する

func main() {
	// panic("runtime error")
	// fmt.Println("Start")

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic("runtime error")
	fmt.Println("Start")
}
