package main

import "fmt"

func main() {

	x := 10
	if x > 10 {
		fmt.Println("x is greater then 5")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is smaller then 5")
	}
}
