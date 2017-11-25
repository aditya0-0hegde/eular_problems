package main

import (
	"fmt"
	"math"
)

func prime_fractors(n int) []int {
	factors := []int{}
	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}
	if n > 2 {
		factors = append(factors, n)
	}
	return factors
}

func main() {
	fmt.Println(prime_fractors(600851475143))
}
