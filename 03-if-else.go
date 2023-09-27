package main

import "fmt"

func thirdMain() {
	age := 25

	if age < 25 {
		fmt.Println("You are not allowed to the party")
	}else if age == 25{
		fmt.Println("You are a lucky. Welcome to the party")
	}else{
		fmt.Println("You are allowed into the party")
	}
}