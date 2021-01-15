package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)

	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 与えた要素数を自動でカウントして変数にセット
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)

	fmt.Println(len(arr1))
}
