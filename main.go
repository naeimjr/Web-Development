package main

import "fmt"

func main() {

	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)

	}(5, 7)

}
