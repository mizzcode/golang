package main 

func getFullName() (firstName, lastName string) {
	firstName = "John"
    lastName = "Doe"
    return
}

func main() {
firstName, lastName := getFullName()
println(firstName, lastName)
}