package main

import "fmt"

type User struct {
	Name        string
	Age         int
	Email       string
	Sex         string
	PhoneNumber string
}

func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
	fmt.Println("Email:", usr.Email)
	fmt.Println("Sex:", usr.Sex)
	fmt.Println("PhoneNumber:", usr.PhoneNumber)
}

func (usr User) printDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
	fmt.Println("Email:", usr.Email)
	fmt.Println("Sex:", usr.Sex)
	fmt.Println("PhoneNumber:", usr.PhoneNumber)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {
	var user1 User
	user1 = User{
		Name:        "Naeim",
		Age:         24,
		Email:       "dm.naeim10@gmail.com",
		Sex:         "Male",
		PhoneNumber: "01736821626",
	}

	//printUserDetails(user1)
	user1.printDetails()
	user1.call(10)

	user2 := User{
		Name:        "Minhaj",
		Age:         30,
		Email:       "dm.minhajsadik@gmail.com",
		Sex:         "Male",
		PhoneNumber: "01781583107",
	}

	//printUserDetails(user2)
	user2.printDetails()
	user2.call(15)

}
