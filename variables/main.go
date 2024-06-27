package main

import "fmt"

func main() {
	my_variable := test()
	fmt.Println(my_variable)
}

func test() string {
	return "testing return"
}

