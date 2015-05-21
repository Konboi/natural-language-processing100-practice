package main

import (
	"fmt"
)

func main() {
	str := "stressed"

	fmt.Printf("before: %s \n", str)

	reverseStr := Reverse(str)

	fmt.Printf("result:  %s \n", reverseStr)

}

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
