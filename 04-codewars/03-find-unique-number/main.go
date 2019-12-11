package main

import "fmt"

func contains(s []float32, e float32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func findUniq(arr []float32) float32 {
	o, t := []float32{arr[0]}, []float32{}

	for x := 1; x < len(arr); x++ {
		i := arr[x]
		if contains(o, i) {
			o = append(o, i)
		} else {
			t = append(t, i)
		}

		// Breaks out of the loop early if it can
		oLen, tLen := len(o), len(t)
		if (oLen > 1 && tLen == 1) || (oLen == 1 && tLen > 1) {
			break
		}
	}

	if len(o) == 1 {
		return o[0]
	}

	return t[0]
}

func main() {
	var test1 = []float32{1, 1, 1, 2, 1, 1}
	var test2 = []float32{0, 0, 0.55, 0, 0}

	// test1 = 2
	fmt.Printf("%f\n", findUniq(test1))

	// test2 = 0.55
	fmt.Printf("%f\n", findUniq(test2))
}

/*

# Best Solution

func FindUniq(arr []float32) float32 {
	// Handles a 3 long arr
	if arr[0] != arr[1] && arr[1] == arr[2] { return arr[0] }

	// Loops through and check the one before
	// if it is different, return the current loop item
  for i, v := range arr[1:] {
    if v != arr[i] { return v }
	}

	// fallback
  return 0
}

*/
