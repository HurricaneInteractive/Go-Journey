package main

import (
	"fmt"
	"math"
)

func century(year int) (c int) {
	c = int(math.Ceil(float64(year) / 100))
	return
}

func main() {
	examples := [][]int{{1990, 20}, {1705, 18}, {1900, 19}}

	for _, e := range examples {
		fmt.Printf("Expect: %d. Got -> %d\n", e[1], century(e[0]))
	}
}
