package main

import "fmt"

func main() {
	numbers :=[]int {100,200,300}
	numbers = append(numbers, doubledigits())
    for i, number := range numbers {
	fmt.Println(i, number)}
}

func doubledigits() int {
	return 10
}