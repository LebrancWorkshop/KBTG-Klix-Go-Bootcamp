package fizzbuzz

import "fmt"

func FizzBuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}

	if n == 5 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", n)

}

