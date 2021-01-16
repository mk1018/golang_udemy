package main

import (
	"fmt"
	"time"
)

// 並行処理

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()

	for {
		fmt.Println("Main")
		time.Sleep(200 * time.Millisecond)
	}
}
