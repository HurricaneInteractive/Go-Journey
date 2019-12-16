package main

import (
	"bytes"
	"strings"
)

func charPosition(s byte) int {
	arr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return bytes.IndexByte([]byte(arr), s)
}

// PrizeDraw method
func PrizeDraw(st []string, we []int, n int) string {
	if len(st) == 0 {
		return "No participants"
	}
	if n > len(st) {
		return "Not enough participants"
	}

	for i, n := range st {
		som := len(n)
		for _, s := range strings.ToUpper(n) {
			som += charPosition(byte(s))
		}
		som = som * we[i]
	}

	return ""
}

func main() {

}
