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
		// log.Println(buf)

		// 先頭末尾の `{{` `}}` を除去
		buf = strings.Trim(buf, "{}")
		// fmt.Println(buf)
		for i, kvStr := range sRe.Split(buf, -1) {
			fmt.Println("-----------------------------")
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
				v = strings.Join(kv[1:], "=") // = が複数回出てくる可能性があるので join する
				// ''でjoinになっていたため、途中の「=」が消えてしまうバグがあったので、「=」でjoinするように修正
				// これで「url=」とか「title=」とかがちゃんと生き残る
			}
			// fmt.Printf("%q\n",kv)
			// 先頭末尾の空白と改行を除去
			k = strings.Trim(k, " \n")
			v = strings.Trim(v, " \n")

			// '' ''' '''' 除去
			v = strongRe.ReplaceAllString(v, "")
			removeMarkup := ""
			// 内部リンクの除去
			removeMarkup = removeInternalLink(v)
			// htmlタグ除去
			removeMarkup = removeHtmlTag(removeMarkup)
			// ファイルを除去
			removeMarkup = removeFile(removeMarkup)
			// 外部リンクを除去
			removeMarkup = removeExternalLink(removeMarkup)
			// Citeを除去
			removeMarkup = removeCiteLink(removeMarkup)

			fmt.Printf("origin key: %s, content: %s \n\n", k, v)
			if removeMarkup == "" {
				fmt.Printf("key: %s, content: %s\n", k, v)
			} else {
				fmt.Printf("key: %s, content: %s\n", k, removeMarkup)
			}
		}
		fmt.Println("=end============================================\n")
	}
}

func removeInternalLink (str string) string {
	var lRe = regexp.MustCompile(`\[\[(.+?\|?.+?)\]\]`)
	findDisplayName := func(s string) string {
		// (.+?\|?.+?)の部分を取得
		subStrs := lRe.FindStringSubmatch(s)
		fileRe := regexp.MustCompile(`^ファイル:`)
		// マッチしていたら かつ 「ファイル:」で始まっていない
		if (len(subStrs) == 2) && !fileRe.MatchString(subStrs[1]) {
			// "|"で分割して最後の要素をreturn
			matchStrs := strings.Split(subStrs[1], "|")
			result := matchStrs[len(matchStrs)-1]
			// もし「#」で始まっていれば「#」を除去
			if strings.HasPrefix(result, "#") {
				tmp := []rune(result)
				result = string(tmp[1:len(tmp)])
			}
			return result
		}
		return s
	}
	// [[val]]の形のやつを検索
	// その中でfindDisplayName関数を実行した返り値に置き換え
	result := lRe.ReplaceAllStringFunc(str, findDisplayName)

	return result
}
func removeHtmlTag (str string) string {
	result := str
	// コメントアウト
	var commentRe = regexp.MustCompile(`<\!--.*?-->`)
	// refタグ
	var refRe = regexp.MustCompile(`<ref.*?>|</ref>`)
	// brタグ（後ろに改行があるときもあるので、それも含めて検索）
	var brRe = regexp.MustCompile(`<br\s?/>\n?`)
	// supタグ
	var supRe = regexp.MustCompile(`<sup>|</sup>`)
	//除去
	result = commentRe.ReplaceAllString(result, "")
	result = refRe.ReplaceAllString(result, "")
	result = brRe.ReplaceAllString(result, "\n")
	result = supRe.ReplaceAllString(result, "")
	return result
}
func removeFile (str string) string {
	var re = regexp.MustCompile(`\[\[(.+)\|(.+)\|(.+)\]\]`)
	findDisplayName := func(s string) string {
		// ((.+)\|(.+)\|(.+))の部分を取得
		subStrs := re.FindStringSubmatch(s)
		if len(subStrs) == 4 {
			return subStrs[len(subStrs)-1]
		}
		return s
	}
	result := re.ReplaceAllStringFunc(str, findDisplayName)
	return result
}
func removeExternalLink (str string) string {
	var re = regexp.MustCompile(`\[(http\S+?)\s(.+?)\]`)
	findDisplayName := func(s string) string {
		linkStrs := re.FindStringSubmatch(s)
		if len(linkStrs) == 2 {
			return "[1]"
		} else if len(linkStrs) == 3 {
			return linkStrs[len(linkStrs)-1]
		}
		return s
	}
	return re.ReplaceAllStringFunc(str, findDisplayName)
}
func removeCiteLink (str string) string {
	var re = regexp.MustCompile(`{{Cite.+?title=(.+?)\|.+?}}`)
	replaceCiteTitle := func(s string) string {
		subStrs := re.FindStringSubmatch(s)
		if (len(subStrs) == 2) {
			return "(" + subStrs[1] + ")"
		}
		return s
	}
	result := re.ReplaceAllStringFunc(str, replaceCiteTitle)
	return result
}
