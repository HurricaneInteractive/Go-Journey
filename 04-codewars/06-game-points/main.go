package main

import (
	"fmt"
	"strings"
)

// Points method
func Points(games []string) (r int) {
	for _, p := range games {
		s := strings.Split(p, ":")
		x, y := s[0], s[1]
		if x > y {
			r += 3
		} else if x == y {
			r++
		}
	}

	return
}

func main() {
	// 30, 10, 0
	examples := [][]string{
		{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"},
		{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"},
		{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"},
	}

	for _, e := range examples {
		fmt.Printf("%d\n", Points(e))
	}
}

/*
// -- Way to do it with only "fmt"

var x, y int

// in loop
fmt.Sscanf(e, "%d:%d", &x, &y)

*/
