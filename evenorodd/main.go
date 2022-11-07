package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range numbers {
		fmt.Println(numbers[i])
		if checkEvenOdd(numbers[i]) {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}

func checkEvenOdd(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
