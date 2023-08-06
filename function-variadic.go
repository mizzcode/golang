package main

func sumAll(numbers ...int) int {
	sum := 0
    for _, n := range numbers {
        sum += n
    }
    return sum
}

func main() {
	sum := sumAll(1,2)
	println(sum)

	// jika punya slice
	slice := []int{1,2,3}
	sum = sumAll(slice...)
	println(sum)
}