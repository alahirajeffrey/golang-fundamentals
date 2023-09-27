package main

import "fmt"

func seventhMain() {

	type Person struct {
		id     int
		age    int
		height string
		gender bool
	}

	jeffrey := Person{
		id:     1,
		age:    25,
		height: "5'7",
		gender: true,
	}

	fmt.Println(jeffrey.age, jeffrey.id, jeffrey.height, jeffrey.gender)

	alahira := Person{2,40,"6'0",true}

	fmt.Println(alahira)
}