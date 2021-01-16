package main

import (
	"fmt"
	"os"
)

func main() {
	TestDefar()

	defer func() {
		fmt.Println("1")
		fmt.Println("2")
	}()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}

func TestDefar() {
	defer fmt.Println("END") // TestDeferが完了したときに実行される
	fmt.Println("START")
}
