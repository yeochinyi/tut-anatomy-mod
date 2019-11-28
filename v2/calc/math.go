package calc

import (
	"errors"
	"fmt"
)

// Add ... returns sum of two integers
func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("Need > 2 number")
	} else {
		for _, num := range numbers {
			sum = sum + num
		}

		fmt.Println("debug info")

		return sum, nil

	}

}
