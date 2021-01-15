package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "1"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
