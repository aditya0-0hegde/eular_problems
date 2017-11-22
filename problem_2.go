package main

import "fmt"

func fibonacci(i int) int {
	if i == 1 {
		return 1
	}
	if i == 2 {
		return 2
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	sum := 0
	i := 1
	for true {
		if fibonacci(i) > 4000000 {
			break
		}
		if fibonacci(i)%2 == 0 {
			sum += fibonacci(i)
		}
		i++
	}
	fmt.Println(sum)
}
