package main

import "fmt"

func main() {

	// 無名関数
	f := func(x, y int) int {
		return x * y
	}

	i := f(1, 2)
	fmt.Println(i)

	// 無名関数に引数を渡す
	i2 := func(x, y int) int {
		return x * y
	}(1, 2)

	fmt.Println(i2)

	// 関数を返す関数
	a := Returnfunc()
	a()

	// 関数を引数にする関数
	CallFunction(func() {
		fmt.Println("funcsion")
	})

}

func CallFunction(g func()) {
	g()
}

func Returnfunc() func() {
	return func() {
		fmt.Println("func")
	}
}
