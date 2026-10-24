package main

import "fmt"

var (
	a = 10
	b = 50
)

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	p := 50
	q := 89

	add(a, p)
	add(a, b)
	add(a, q)

}
