package calc

import "fmt"

// returns sum of two integers
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum = sum + num
	}

	fmt.Println("debug info")

	return sum
}
