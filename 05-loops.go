package main

import "fmt"

func fifthMain() {
	// first for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("first for loop ended")

	// second for loop
	j := 1
	for j <= 20{
		if j % 2 == 0 {
			j++
			continue
		}
		fmt.Println(j)
		j++
	} 
	fmt.Println(" Second for loop ended")
}