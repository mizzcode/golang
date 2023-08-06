package main

import "fmt"

func main() {
	fmt.Println("Mencari Bilangan Ganjil")
	for i := 0; i <= 10; i++ { 
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Loop -", i)
	}

	fmt.Println("Mencari Bilangan Genap")
	for j := 0; j <= 10; j++ { 
		if j % 2 == 1 {
            continue
        }
        fmt.Println("Loop -", j)
	}
}