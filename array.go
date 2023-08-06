package main

import "fmt"

func main() {

	var names [5]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Mizz"
	names[3] = "Jani"
	names[4] = "Salman"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])
	fmt.Println(names[4])

	var values = [3]int{
		12,
		213,
		123,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))
}