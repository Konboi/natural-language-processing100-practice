package main

import (
	"fmt"
)

func main() {
	str := "パタトクカシーー"
	target := []int{1, 3, 5, 7}
	result := []rune{}
	runes := []rune(str)

	fmt.Printf("before: %s \n", str)

	for _, t := range target {
		targetRune := runes[t-1]
		result = append(result, targetRune)
	}

	fmt.Printf("result: %s \n", string(result))
}
