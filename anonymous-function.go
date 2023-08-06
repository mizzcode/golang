package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println(name, "is blacklisted")
	}else {
		fmt.Println(name, "welcome")
	}
}

func main() {

	registerUser("mizz", func(name string) bool {
		return name == "mizz"
	})

}