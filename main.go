package main

import "fmt"

var (
	a = 10
	b = 20
	x = 80
	y = 69
	e = 31
	w = 78
)

func printNum(num int) {
	fmt.Println(num)
}

func add(a int, b int) {
	res := a + b
	printNum(res)
}
func main() {
	add(a, b)
	add(a, w)
	add(x, x)
	add(e, y)

}
