package main

import (
	"flag"
	"fmt"
	"strings"
)


func main() {
	var (
		n int
		str string
		word bool
	)
	flag.IntVar(&n, "n", 2, "Ngram number")
	flag.StringVar(&str, "s", "面白法人カヤック", "string")
	flag.BoolVar(&word, "w", false, "split space")
	flag.Parse()

	fmt.Printf("%s\n", str)

	if word == true {
		words := strings.Split(str, " ")
		for i := 0; i < len(words); i++ {
			j := i + n
			if j > len(words) {
				j = i
			}
			fmt.Println(words[i:j])
		}
	} else {
		runes := []rune(str)
		for i := 0; i < len(runes); i++ {
			j := i + n
			if j > len(runes) {
				j = i
			}
			fmt.Println(string(runes[i:j]))
		}
	}
}


