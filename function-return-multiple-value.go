package main 

import "fmt"

func getFullName() string, int, string  { 
	return "Mizz", 19, "Kun"
}

func main() { 
	firstName, middleName, lastName := getFullName()
	fmt.Println("Nama:" firstName, "Umur:" middleName, lastName)
}