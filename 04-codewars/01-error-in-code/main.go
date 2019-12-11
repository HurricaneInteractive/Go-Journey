/**
The code does not execute properly. Try to figure out why.

Initial code:
package multiply

func Multiply(a, b int) int {
	a * b
}
*/

package main

import "fmt"

// Multiply Method
func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Printf("1x1 = 1. Returns %d\n", Multiply(1, 1))
	fmt.Printf("1x0 = 0. Returns %d\n", Multiply(1, 0))
	fmt.Printf("2x5 = 10. Returns %d\n", Multiply(2, 5))
	fmt.Printf("5x10 = 50. Returns %d\n", Multiply(5, 10))
}
