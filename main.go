package main

import "fmt"

var arr2 = [3]string{"i", "love", "you"}

func main() {

	var arr [4]int

	arr[3] = 9
	arr[2] = 6
	arr[1] = 3

	fmt.Println(arr)
	fmt.Println(arr2)
}
