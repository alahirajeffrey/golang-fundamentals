package main

import "fmt"

func ninthMain() {
	englishScores := []int{60, 70, 80, 90}

	fmt.Println(englishScores)

	englishScores = append(englishScores, 40,45)

	fmt.Println(englishScores)

	firstScoreSlice := englishScores[0:1]
	secondScoreSlice := englishScores[5:]

	englishSlice := append(firstScoreSlice, secondScoreSlice...)

	fmt.Println(englishSlice)
}