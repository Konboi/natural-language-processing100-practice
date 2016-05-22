package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// Article
type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// 大体以下みたいな形してるので頑張る
// ```
// {{基礎情報 国
// |key = valvalval
// |key = valvalval
// |key = val
// val
// val
// }}
// ```
// ```
// {{基礎情報 国|
// key = valvalval|
// key = valvalval|
// key = val
// val
// val
// }}
// ```
var sRe = regexp.MustCompile(`\n\||\|\n`)

// '' ''' ''''
var strongRe = regexp.MustCompile(`'{2,4}`)

func main() {
	raw, err := ioutil.ReadFile("../20/uk.json")
	if err != nil {
		log.Println(err)
	}

	var articles []Article
	if err := json.Unmarshal(raw, &articles); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, article := range articles {
		// `{{基礎情報 ... }}` を抽出
		var (
			indent = 0
			flag   = false
			buf    = ""
			rv     = []rune(article.Text)
		)
		for i := 0; i < len(rv); i++ {
			r := rv[i]
			switch r {
			case '{':
				if flag == true && "{{" == string(rv[i:i+2]) {
					indent++
				}
				if flag == false && "{{基礎情報" == string(rv[i:i+6]) {
					flag = true
				}
			case '}':
				if flag == true && indent == 0 && "}}" == string(rv[i:i+2]) {
					buf += "}}"
					flag = false
					break
				}
				if flag == true && "}}" == string(rv[i:i+2]) {
					indent--
				}
			}
			if flag == true {
				buf += string(r)
			}
		}
		//log.Println(buf)

		// 先頭末尾の `{{` `}}` を除去
		buf = strings.Trim(buf, "{}")

		for i, kvStr := range sRe.Split(buf, -1) {
			if i == 0 {
				continue
			}
			if kvStr == "" {
				continue
			}

			kv := strings.Split(kvStr, "=")
			k := kv[0]
			v := ""
			if 2 <= len(kv) {
				v = strings.Join(kv[1:], "") // = が複数回出てくる可能性があるので join する
			}

			// 先頭末尾の空白と改行を除去
			k = strings.Trim(k, " \n")
			v = strings.Trim(v, " \n")

			// '' ''' '''' 除去
			v = strongRe.ReplaceAllString(v, "")
			// 内部リンクの除去
			// [[]] 除去
			removeInternalLinkStr := removeInternalLink(v)

			fmt.Printf("origin key: %s, content: %s \n", k, v)
			if removeInternalLinkStr == "" {
				fmt.Printf("key: %s, content: %s\n", k, v)
			} else {
				fmt.Printf("key: %s, content: %s\n", k, removeInternalLinkStr)
			}
		}
		fmt.Println("--------------------------")
	}
}

func removeInternalLink (str string) string {
	var lRe = regexp.MustCompilePOSIX(`\[\[.+?\|+?.+?\]\].+?`)
	var linkRe   = regexp.MustCompile(`[\[|\]]`)
	result := ""
	for _, linkStr := range lRe.FindAllString(str, -1) {
		fmt.Printf("have link: %s\n", linkStr)
		tmpResult := []string{}
		linkVals := strings.Split(linkStr, "|")
		for i, linkChar := range linkVals {
			if i == (len(linkVals) - 1) {
				// |で分割した最後は表示名なので結果に加える
				tmpResult = append(tmpResult, linkChar)
			} else {
				if strings.Contains(linkChar, "[[") {
					subChar := strings.Split(linkChar, "[[")
					for i, char := range subChar {
						if i < (len(subChar) - 1) {
							tmpResult = append(tmpResult, char)
						}
					}
				}
			}
		}
		linkStr = strings.Join(tmpResult, "")
		linkStr = linkRe.ReplaceAllString(linkStr,"")
		result = linkStr
	}
	return result
}
