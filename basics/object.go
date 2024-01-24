package basics

import (
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Address struct {
		Pincode int
		City    string
		State   string
	}
	Mobile int
}

func PrintObjects() {
	user2 := User{
		Name: "Piyush",
		Age:  0,
		Address: struct {
			Pincode int
			City    string
			State   string
		}{Pincode: 843434, City: "GGN", State: "haryana"},
		Mobile: 1234567890,
	}

	fmt.Println(user2)
}
