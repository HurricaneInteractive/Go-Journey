package main

import "fmt"

// TwoSum method
func TwoSum(numbers []int, target int) [2]int {
	if len(numbers) == 2 {
		return [2]int{0, 1}
	}

	len := len(numbers)

	for i := 0; i < len; i++ {
		for x := 1; x < len; x++ {
			if numbers[i]+numbers[x] == target {
				return [2]int{i, x}
			}
		}
	}

	return [2]int{}
}

func main() {
	// ans := TwoSum([]int{1, 3}, 4) // {0,1}
	// ans := TwoSum([]int{1, 2, 3}, 4) // {0,2}
	ans := TwoSum([]int{1234, 5678, 9012}, 14690) // {1,2}

	fmt.Printf("{%d, %d}\n", ans[0], ans[1])
}
