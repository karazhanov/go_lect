package funcs

import "fmt"

func Swap1() {
	var a = 1
	var b = 2

	fmt.Println("a = ", a, " b = ", b)

	tmp := a
	a = b
	b = tmp

	fmt.Println("a = ", a, " b = ", b)
}























func Swap2() {
	var a = 1
	var b = 2

	fmt.Println("a = ", a, " b = ", b)

	a, b, _ = b, a, 3

	fmt.Println("a = ", a, " b = ", b)
}

















func Swap3(a, b int) (int, int) {
	return b, a
}