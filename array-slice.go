package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1) // slice reference ke array

	// slice1[0] = "Bibi May"
	// fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Mizz")
	fmt.Println(slice3)

	slice3[1] = "Jani"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 3, 5)

	newSlice[0] = "Mizz"
	newSlice[1] = "Jani"
	newSlice[2] = "Ferdi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	iniArray := [...]string {
		"Jan", "Feb",
	}
	iniSlice := []string {
		"Jan", "Feb",
	}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}