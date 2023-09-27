package main

import "fmt"

// check if student passed course
func isCoursePassed(score int){
	if (score < 40){
		fmt.Println("course failed")
	}else{
		fmt.Println("course passed")
	}
}

func eigthMain() {
	// arrays
	mathScores := [6]int{10, 20, 30, 40, 50, 70}
	// check if student passed
	for i:= 0; i < len(mathScores); i++{
		isCoursePassed(mathScores[i])
	}
}