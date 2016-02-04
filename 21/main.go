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

var re = regexp.MustCompilePOSIX(`^\[\[Category:.+?\]\].*?`)

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
		for _, str := range re.FindAllString(a.Text, -1) {
			fmt.Println(str)
		}
	}
}
