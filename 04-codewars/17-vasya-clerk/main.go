package main

import (
	"fmt"
)

// Tickets Method
func Tickets(peopleInLine []int) string {
	change := map[int]int{
		25:  0,
		50:  0,
		100: 0,
	}

	for _, p := range peopleInLine {
		switch p {
		case 25:
			change[p]++
			break
		case 50:
			if change[25] < 1 {
				return "NO"
			}
			change[25]--
			change[p]++
			break
		case 100:
			if change[50] >= 1 && change[25] >= 1 {
				change[50]--
				change[25]--
			} else if change[25] >= 3 {
				change[25] -= 3
			} else {
				return "NO"
			}

			change[p]++
			break
		}
	}

	return "YES"
}

func main() {
	ans := []string{"YES", "NO", "NO"}
	examples := [][]int{{25, 25, 50}, {25, 100}, {25, 25, 50, 50, 100}}

	for i, e := range examples {
		fmt.Printf("Expecting: %s. Got -> %s\n", ans[i], Tickets(e))
	}
}

/*

Another solution. Negative check instead of specific checks

func Tickets(peopleInLine []int) string {
  t, f := 0, 0
  for _, bill := range peopleInLine {
    switch bill {
    case 25:
      t++
    case 50:
      f++
      t--
    case 100:
      if f > 0 {
        f--
        t--
      } else {
        t -= 3
      }
    }
    if t < 0 || f < 0 { return "NO" }
  }
  return "YES"
}

*/
