package main

import (
	"fmt"
)

func main() {
	person := map[string]string {
		"name": "Mizz",
		"age":  "19",
	}

	fmt.Println(person)
	fmt.Println("Nama: ", person["name"])
	fmt.Println("Umur:", person["age"])
	fmt.Println("Panjang Map:", len(person))

	books := make(map[string]string) 
		books["book1"] = "Book 1"
        books["book2"] = "Book 2"
        books["book3"] = "Book 3"
        books["ups"] = "Salah"

	fmt.Println(books)
	fmt.Println("Panjang Book:", len(books))

	delete(books, "ups")

	fmt.Println(books)
	fmt.Println("Panjang Book:", len(books))
}