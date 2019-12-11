package main

import "fmt"

// Summation returns the sum from 1 -> n
func Summation(n int) (s int) {
	if n == 1 {
		return n
	}
	s = 0

	for i := 1; i <= n; i++ {
		s = s + i
	}

	return
}

func main() {
	examples := [...][2]int{
		{1, 1},
		{8, 36},
		{22, 253},
		{100, 5050},
		{213, 22791},
	}

	for _, i := range examples {
		fmt.Printf("%d should be: %d. Got -> %d\n", i[0], i[1], Summation(i[0]))
	}
}
