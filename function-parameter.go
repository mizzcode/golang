package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Halo!", firstName, lastName)
}

func main() {

	firstName := "Mizz"
	lastName := "Kun"

	sayHello(firstName, lastName)


	sayHello("Jani", "Chan")

}