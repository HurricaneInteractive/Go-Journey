package main

import (
	"fmt"
	"sort"
	"strings"
)

// TwoToOne Method
func TwoToOne(s1 string, s2 string) string {
	s := make(map[string]bool)
	rtn := []string{}

	for _, a := range s1 + s2 {
		if _, ok := s[string(a)]; !ok {
			s[string(a)] = true
			rtn = append(rtn, string(a))
		}
	}

	sort.Strings(rtn)
	return strings.Join(rtn, "")
}

func main() {
	fmt.Printf("Expecting: aehrsty. Got -> %s\n", TwoToOne("aretheyhere", "yestheyarehere"))
	fmt.Printf("Expecting: abcdefghilnoprstu. Got -> %s\n", TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
}
