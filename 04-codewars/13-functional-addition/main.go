package main

// Add method
func Add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
