package main

import "fmt"

// GetSize of a box
func GetSize(w, h, d int) [2]int {
	v := w * h * d
	a := ((d * w) * 2) + ((d * h) * 2) + ((w * h) * 2)

	return [2]int{a, v}
}

func main() {
	examples := [...][5]int{{4, 2, 6, 88, 48}, {1, 1, 1, 6, 1}, {1, 2, 1, 10, 2}, {1, 2, 2, 16, 4}}

	for _, i := range examples {
		size := GetSize(i[0], i[1], i[2])
		fmt.Printf("Expecting: %d and %d\n", i[3], i[4])
		fmt.Printf("Area: %d, Volumn: %d\n\n", size[0], size[1])
	}
}
