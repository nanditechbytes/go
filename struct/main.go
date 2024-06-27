package main

import "fmt"

type Person struct {
	name string
	age int
	seat int
}

func main() {
	person1 := Person{
		name: "John",
		age: 20,
		seat: 1,
	}

	var person2 Person
	person2 = Person{
		name: "Paul",
		age: 25,
		seat: 5,
	}

	fmt.Println("Passenger1", person1)
	fmt.Println("Passeneger2", person2.name)
	fmt.Println("Passenger1_name", person1.name)
	fmt.Println("Passenger1_age", person1.age)
}