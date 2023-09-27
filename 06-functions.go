package main

import "fmt"

func countDown(firstNum int, secondNum int){
	fmt.Println("count down started")
	for i := firstNum; i >= secondNum; i-- {
		fmt.Println(i)
	}
}

func sixthMain(){
	countDown(10, 1)
}