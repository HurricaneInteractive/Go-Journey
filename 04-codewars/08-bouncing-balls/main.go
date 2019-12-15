package main

import "fmt"

// BouncingBall Method
func BouncingBall(h, bounce, window float64) (c int) {
	if h <= 0 || (bounce <= 0 || bounce >= 1) || window >= h {
		return -1
	}

	for h >= window {
		h = h * bounce
		// bounce = h / bounce
		c++
	}

	return c + 1
}

func main() {
	examples := [][]float64{{3, 0.66, 1.5, 3}, {40, 0.4, 10, 3}, {10, 0.6, 10, -1}, {40, 1, 10, -1}}
	// examples := [][]float64{{3, 0.66, 1.5, 3}}
	// examples := [][]float64{{40, 0.4, 10, 3}}

	for _, e := range examples {
		fmt.Printf("Expect %f. Got -> %d\n", e[3], BouncingBall(e[0], e[1], e[2]))
	}
}
