package main

import "fmt"

func getHello(fullname string) string { 
	if fullname == "" {
		return "Hello MasBro"
	} else {
		return "Hello " + fullname
	}
}

func getZero(fullname string) {
	if fullname == "" {
        fmt.Println("zero")
    } else {
        fmt.Println("zero " + fullname)
    }
}

func main() {

	name := getHello("Mizz")

	fmt.Println(name)

	name = getHello("")
	fmt.Println(name)

	// jika func tidak mengembalikan value maka error
	// tes := getZero("Jani")
	// fmt.Println(tes)

}