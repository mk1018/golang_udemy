package main

import "fmt"

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("a")
	default:
		fmt.Println("b")
	}

	switch i := 2; i {
	case 1, 2:
		fmt.Println("a")
	}

	f := 6
	switch {
	case f > 0 && f < 4:
		fmt.Println("a")
	case f > 3:
		fmt.Println("b")
	}

	type_switch()
}

func type_switch() {
	anything("aaa")

	var x interface{} = "aiueo"

	i, isInt := x.(int)
	fmt.Println(i, isInt)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("a")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "is String")
	}

	switch x.(type) {
	case int:
		fmt.Println("a")
	case string:
		fmt.Println("b")
	}
}

func anything(a interface{}) {
	fmt.Println(a)
}
