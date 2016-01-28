package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	checkString = "イギリス"
)

func main() {
	checker := regexp.MustCompile(checkString)

	filename := "jawiki-country.json.gz"
	file, err := os.Open(filename)

	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer gz.Close()

	r := bufio.NewReader(gz)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Println(err.Error())
			return
		}
		if checker.FindString(line) == "" && checker.FindString(line) == "" {
			continue
		}

		article := &Article{}
		if err = json.Unmarshal([]byte(line), article); err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Println(article.Text)
	}

}
