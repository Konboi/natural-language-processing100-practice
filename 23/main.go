package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

var re = regexp.MustCompilePOSIX(`^(=.+?)$`)
var s1 = regexp.MustCompilePOSIX(`^=[^=]+`)
var s2 = regexp.MustCompilePOSIX(`^==[^=]+`)
var s3 = regexp.MustCompilePOSIX(`^===[^=]+`)
var s4 = regexp.MustCompilePOSIX(`^====[^=]+`)
var s5 = regexp.MustCompilePOSIX(`^====[^=]+`)

func main() {
	raw, err := ioutil.ReadFile("../20/uk.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var article []Article
	if err := json.Unmarshal(raw, &article); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, a := range article {
		for _, str := range re.FindAllStringSubmatch(a.Text, -1) {
			var n int
			if s1.MatchString(str[1]) {
				n = 1
			}
			if s2.MatchString(str[1]) {
				n = 2
			}
			if s3.MatchString(str[1]) {
				n = 3
			}
			if s4.MatchString(str[1]) {
				n = 4
			}
			if s5.MatchString(str[1]) {
				n = 5
			}
			var rep = regexp.MustCompile(`=`)
			var section = rep.ReplaceAllString(str[1], "")
			fmt.Printf("%d:%s\n", n, section)
		}
	}
}
