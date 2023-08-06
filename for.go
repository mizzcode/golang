package main

import "fmt"

func main() {
	// counter := 1
	
	// for counter <= 10 {
	// 	fmt.Println(counter)
    //     counter++
	// }

	for counter := 1; counter <= 10; counter++ { 
		fmt.Println("Test", counter, "Passed")
	}

	slice := []string {
		"Mizz",
		"Kun",
		"Jani",
		"Chan",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice { 
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Mizz Kun"
	person["age"] = "19"
	person["gender"] = "Male"

	for key, value := range person { 
		fmt.Println(key, "=", value)
	}
}