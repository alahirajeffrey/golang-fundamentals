package main

import "fmt"

func tenthMain() {
	// create map i.e key value pair
	cart := make(map[string]int)

	// write to cart
	cart["food"] = 100
	cart["beverages"] = 50
	cart["drugs"] = 20

	fmt.Println(cart)

	// read from map
	foodQuantity, isFound := cart["food"]

	fmt.Println(foodQuantity, isFound)

	// delete from map
	delete(cart, "drugs")

	fmt.Println(cart)

}