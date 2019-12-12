package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Points method
func Points(games []string) (r int) {
	for _, p := range games {
		s := strings.Split(p, ":")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		if x > y {
			r += 3
		}
		if x == y {
			r++
		}
	}

	return
}

func main() {
	// 30, 10, 0
	examples := [...][]string{
		{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"},
		{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"},
		{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"},
	}

	for _, e := range examples {
		fmt.Printf("%d\n", Points(e))
	}
}
