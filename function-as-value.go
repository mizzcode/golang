package main

import "fmt"

func getName(name string) string {
	return "Hi " + name
}

func main() {
	sayHi := getName

	result := sayHi("Mizz")
	fmt.Println(result)

	// sebenarnya sama seperti
	fmt.Println(getName("Jani"))
}