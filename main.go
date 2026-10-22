package main

import "fmt"

var (
	a  = 20
	b  = 20
	x  = 570
	y  = 69
	e  = 31
	w  = 999
	xy = 101
)

func printNum(num int) {
	fmt.Println(num)
}

func add(a int, b int) {
	res := a // xy
	printNum(res)
}
func main() {
	add(a, b)
	add(a, w)
	add(x, x)
	add(e, xy)

}
