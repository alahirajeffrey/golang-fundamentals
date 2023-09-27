package main

import (
	"fmt"
)

func firstMain() {

	// variables
	var name string
	var age int
	var isAdult bool

	// compound variable creation
	name, age, isAdult= "Jeffrey Alahira", 17, false

	fmt.Println(name, age, isAdult)

	// block variable creation
	var (
		isParent = false
		isStudent = true
		dateOfBirth = "31st september 2020"
	
	)	

	fmt.Println(isParent, isStudent, dateOfBirth)

	// constants
	 pi := 3.14

	fmt.Println(pi)

}