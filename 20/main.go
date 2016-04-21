package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Articles []*Article

var (
	checkString = "イギリス"
)

func main() {
	checker := regexp.MustCompile(checkString)

	filename := "jawiki-country.json.gz"
	file, err := os.Open(filename)
	var articles Articles

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
			if err.Error() == "EOF" {
				break
			}

			log.Println("readline :", err.Error())
			return
		}
		if checker.FindString(line) == "" {
			continue
		}

		article := &Article{}
		if err = json.Unmarshal([]byte(line), article); err != nil {
			log.Println(err.Error())
		}

		articles = append(articles, article)

		fmt.Println(article.Text)
	}

	jsonFileNmae := "uk.json"
	jsonData, err := json.Marshal(articles)
	if err != nil {
		log.Println("marshal error:", err.Error())
		return
	}
	err = ioutil.WriteFile(jsonFileNmae, jsonData, 0644)
	if err != nil {
		log.Println("publish error:", err.Error())
		return
	}
}
