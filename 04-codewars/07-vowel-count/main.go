package main

import (
	"fmt"
	"regexp"
)

// GetCount method
func GetCount(str string) int {
	re := regexp.MustCompile("[^a|e|i|o|u]")
	return len(re.ReplaceAllString(str, ""))
}

func main() {
	// 5
	examples := []string{"abracadabra"}

	for _, e := range examples {
		fmt.Printf("%d\n", GetCount(e))
	}
}
