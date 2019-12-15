package main

import "fmt"

// FindOdd Method
func FindOdd(seq []int) int {
	acc := make(map[int]int)

	if len(seq) == 1 {
		return seq[0]
	}

	for _, x := range seq {
		acc[x]++
	}
	for i, a := range acc {
		if a%2 != 0 {
			return i
		}
	}

	return 0
}

func main() {
	arr := [...][]int{
		{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
		{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
		{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
		{10},
		{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
		{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
	}
	sol := [...]int{5, -1, 5, 10, 10, 1}

	for i, v := range arr {
		fmt.Printf("Expecting: %d. Got -> %d\n", sol[i], FindOdd(v))
	}
}
