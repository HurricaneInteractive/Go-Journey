package main

import (
	"fmt"
	"regexp"
)

func errorPrinter(s string) int {
	re := regexp.MustCompile(`[a-m]`)
	s = re.ReplaceAllString(s, "")
	return len(s)
}

func main() {
	// 0/14, 8/22, 3/56
	test := [3]string{"aaabbbbhaijjjm", "aaaxbbbbyyhwawiwjjjwwm", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"}

	for _, s := range test {
		fmt.Printf("%d/%d\n", errorPrinter(s), len(s))
	}
}
