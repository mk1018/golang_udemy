package main

import "fmt"

func main() {
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
		}
	}
}
