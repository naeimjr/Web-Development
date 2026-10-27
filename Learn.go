package main

import "fmt"

const a = 100

var p = 99

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}
	add(5, 6)
	add(p, a)
}
func main() {
	call()
	fmt.Println(a)
}
func init() {
	fmt.Println("hello naeim")
}
