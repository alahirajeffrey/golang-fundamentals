package main

import "fmt"

func secondMain() {
	var firstNumber float32 = 15.0
	var secondNumber float32 = 3.28

	// arithmetic oprators
	addition := firstNumber + secondNumber
	subtraction := firstNumber - secondNumber
	division := firstNumber / secondNumber
	multiplication := firstNumber * secondNumber

	fmt.Println(addition, subtraction, division, multiplication)

	// relational operators
	firstResult := firstNumber >= secondNumber
	secondResult := firstNumber <= secondNumber
	thirdResult := firstNumber != secondNumber

	fmt.Println(firstResult, secondResult, thirdResult)

}