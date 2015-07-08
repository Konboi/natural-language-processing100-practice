package main

import (
	"fmt"
	"regexp"
)

var (
	input = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	re    = regexp.MustCompile(`\.?(\s+|$)`)
	words = re.Split(input, -1)
)

func main() {
	res := make(map[string]int)

	for i, v := range words {
		if v == "" {
			break
		}

		pos, r := i+1, []rune(v)

		switch pos {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19:
			res[string(r[:1])] = pos
		default:
			res[string(r[:2])] = pos
		}
	}

	for k, v := range res {
		fmt.Printf("%d\t%s\n", v, k)
	}
}
