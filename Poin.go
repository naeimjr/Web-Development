package main

import "fmt"

//func
// print(numbers *[3]int) {
//fmt.Println(numbers)
//}

func main() {

	arr := [6]string{"This", "is", "a", "Go", "Interview"}
	fmt.Println(arr)

	s := arr[1:4]
	fmt.Println(s)

	s1 := s[1:2]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	//x := 20
	//fmt.Println("x =", x)
	//p := &x

	//*p = 30
	//fmt.Println("x =", x)
	//fmt.Println("Address:", p)
	//fmt.Println("Value at the address:", *p)
	//arr := [3]int{1, 2, 3}
	//print(&arr)
}
