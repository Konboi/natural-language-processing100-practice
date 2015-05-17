package main

import (
	"fmt"
)

func main() {
	str := "パタトクカシーー"
	target := []int{1, 3, 5, 7}
	result := []rune{}

	for _, v := range target {
		targetRune := GetTargetRune(str, v-1)
		result = append(result, targetRune)
	}

	fmt.Println(string(result))

}

func GetTargetRune(s string, target int) rune {
	runes := []rune(s)

	return runes[target]
}
