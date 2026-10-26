package main

import "fmt"

type User struct {
	Name        string
	Age         int
	Email       string
	Sex         string
	PhoneNumber string
}

func main() {
	var user1 User
	user1 = User{
		Name:        "Naeim\t \"Minhaj\"",
		Age:         24,
		Email:       "dm.naeim10@gmail.com",
		Sex:         "Male",
		PhoneNumber: "01736821626",
	}

	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)
	fmt.Println("Email:", user1.Email)
	fmt.Println("Sex:", user1.Sex)
	fmt.Println("PhoneNumber:", user1.PhoneNumber)
	fmt.Print("\n")
	user2 := User{
		Name:        "Minhaj\tNaeim",
		Age:         30,
		Email:       "dm.minhajsadik@gmail.com",
		Sex:         "Male",
		PhoneNumber: "01781583107",
	}
	fmt.Println("Name:", user2.Name)
	fmt.Println("Age:", user2.Age)
	fmt.Println("Email:", user2.Email)
	fmt.Println("Sex:", user2.Sex)
	fmt.Println("PhoneNumber:", user2.PhoneNumber)
}
