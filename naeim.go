package main

import "fmt"

func main() {
	arr := [3]string{"a", "b", "c"} // array: always exactly 3
	fmt.Println(arr, len(arr))

	var ids []string // slice: empty for now
	fmt.Println(ids, len(ids))

	ids = append(ids, "x")
	ids = append(ids, "y")
	fmt.Println(ids, len(ids))

	for i, id := range ids {
		fmt.Println(i, id)
	}
}
