package main

// SumCubes method
func SumCubes(n int) (s int) {
	for i := 1; i <= n; i++ {
		s += i * i * i
	}

	return
}

/*
Clever solution

func SumCubes(n int) int {
  n = n * (n + 1)
  return n * n / 4
}
*/
