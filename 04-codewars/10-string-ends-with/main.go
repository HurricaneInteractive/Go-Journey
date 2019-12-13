package main

import (
	"fmt"
	"strings"
)

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func main() {
	examples := [][]string{{"", ""}, {" ", ""}, {"abc", "c"}, {"banana", "ana"}, {"a", "z"}, {"", "t"}}
	ans := []bool{true, true, true, true, false, false}

	for i, e := range examples {
		fmt.Printf("Expected: %t. Got -> %t\n", ans[i], solution(e[0], e[1]))
	}
}
