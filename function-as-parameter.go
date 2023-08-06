package main

import "fmt"

func sayHello(name string, filter func(string)string) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" || name == "Babi" {
		return "***"
	} else {
		return name
	}
}


func main() {
	sayHello("Babi", spamFilter)
}