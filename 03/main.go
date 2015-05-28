package main

import (
	"fmt"
	"strings"
)

/*
03. 円周率
"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
*/

var s = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

func main() {
	counts := CountWordStrings(s)
	fmt.Println(counts)
}

func CountWordStrings(s string) []int {
	counts := []int{}

	count := 0
	for _, r := range s {
		switch r {
		case ' ', ',', '.':
			if count > 0 {
				counts = append(counts, count)
				count = 0
			}
			continue
		default:
			count++
		}
	}
	return counts
}

func CountWordStringsTrim(s string) []int {
	s = strings.Trim(s, ".")            // "."の除去、Trimは文字列の前後しか除去できない
	s = strings.Replace(s, ",", "", -1) //","の除去

	words := strings.Split(s, " ") //slicesに単語が格納

	counts := []int{}
	for _, word := range words {
		counts = append(counts, len(word))
	}

	return counts
}

// http://mattn.kaoriya.net/software/lang/go/20150416205650.htm
func CountWordStringsMattn(s string) []int {
	var pi []int
	var prev, i int
	b := []byte(s)
	l := len(b)
	for i < l {
		for ; i < l; i++ {
			if ('a' <= b[i] && b[i] <= 'z') || ('A' <= b[i] && b[i] <= 'Z') {
				break
			}
		}
		if i < l {
			prev = i
			for ; i < l; i++ {
				if !(('a' <= b[i] && b[i] <= 'z') || ('A' <= b[i] && b[i] <= 'Z')) {
					break
				}
			}
			pi = append(pi, i-prev)
		}
		i++
	}

	return pi
}

// https://gist.github.com/methane/ff65135764556996989b
func CountWordStringsMethane(s string) []int {
	res := make([]int, 0, len(s)/6)
	cnt := 0
	for _, c := range s {
		switch {
		case 'a' <= c && c <= 'z', 'A' <= c && c <= 'Z':
			cnt++
		default:
			if cnt != 0 {
				res = append(res, cnt)
				cnt = 0
			}
		}
	}
	if cnt != 0 {
		res = append(res, cnt)
		cnt = 0
	}
	return res
}
