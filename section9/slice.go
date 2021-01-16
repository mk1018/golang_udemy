package main

import "fmt"

func main() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B", "C", "D"}
	fmt.Println(sl3)
	fmt.Println(sl3[1])

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	fmt.Println(sl3[:3])

	sl5 := map[string]int{"A": 100, "B": 200}
	fmt.Println(sl5)
}
