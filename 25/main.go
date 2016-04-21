package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Article
type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type infoObject struct {
	Key     string
	Content string
}

type infoObjects []infoObject

func main() {
	log.SetFlags(log.Lshortfile)
	raw, err := ioutil.ReadFile("../20/uk.json")

	if err != nil {
		log.Println(err.Error())
	}

	var article []Article
	if err := json.Unmarshal(raw, &article); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var infoObjs infoObjects
	for _, a := range article {
		info := getIndexContent(strings.Replace(a.Text, "\n", "", -1))

		infoStrings := divideInfo(info)
		for _, s := range infoStrings {
			trimString := strings.TrimSpace(s)
			//fmt.Println(trimString)
			splitStrings := strings.SplitN(trimString, "=", 2)
			//fmt.Println(splitStrings)
			if len(splitStrings) == 2 {
				infoObj := infoObject{
					Key:     splitStrings[0],
					Content: splitStrings[1],
				}

				infoObjs = append(infoObjs, infoObj)
			}
		}
		//fmt.Println(string(info))
		//fmt.Println("_______________")
	}

	for _, v := range infoObjs {
		fmt.Printf("key: %s, content: %s \n", v.Key, v.Content)
	}
}

func getIndexContent(article string) string {
	runes := []rune(article)
	contentFlag := false

	var buf []rune
	var bufEnd []rune
	var parse []rune
	var contents []rune
BREAK_CONTENT:
	for _, v := range runes {
		switch string(v) {
		case "{":
			if len(buf) == 0 {
				buf = append(buf, v)
			} else if contentFlag {
				parse = append(parse, v)
			} else {
				continue
			}
		case "基":
			if len(buf) == 1 {
				buf = append(buf, v)
			} else {
				continue
			}
		case "礎":
			if len(buf) == 2 {
				contentFlag = true
			} else {
				continue
			}
		case "}":
			if contentFlag {
				if len(parse) == 0 {
					if len(bufEnd) == 1 {
						contentFlag = false
						break BREAK_CONTENT
					} else {
						bufEnd = append(bufEnd, v)
					}
				} else {
					parse = parse[:len(parse)-1]
				}
			} else {
				continue
			}

		default:
			if contentFlag == true {
				contents = append(contents, v)
			}
		}
	}

	return string(contents)
}

func divideInfo(info string) []string {
	runes := []rune(info)

	var infoStrings []string
	var infoStringRunes []rune

	parseStart := false
	for _, v := range runes {
		switch string(v) {
		case "|":
			if parseStart {
				infoStrings = append(infoStrings, string(infoStringRunes))
				infoStringRunes = []rune{}
				break
			} else {
				parseStart = true
			}
		default:
			if parseStart {
				infoStringRunes = append(infoStringRunes, v)
			}
		}
	}

	return infoStrings
}
