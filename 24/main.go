package main

/*
24. ファイル参照の抽出
記事から参照されているメディアファイルをすべて抜き出せ．
*/

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	SOURCE_FILENAME = "jawiki-country.json.gz"
)

type Data struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func main() {
	var err error

	f, err := os.Open(SOURCE_FILENAME)
	if err != nil {
		log.Fatalf("source file open error: %s", err)
	}

	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatalf("source file gzip read error: %s", err)
	}

	dec := json.NewDecoder(gr)
	for {
		var decoded Data
		if err := dec.Decode(&decoded); err != nil {
			break
		}
		b := bytes.NewBufferString(decoded.Text)
		s := NewWikiScanner(b)
		for {
			link, err := s.LookupLink()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%s", err)
			}

			if link.Type == "File" || link.Type == "ファイル" {
				fmt.Println(link.Name)
			}
		}
	}
}

type WikiScanner struct {
	r io.RuneReader
}

func NewWikiScanner(r io.RuneReader) *WikiScanner {
	return &WikiScanner{r: r}
}

type WikiFileLink struct {
	Type    string
	Name    string
	Options string
	Caption string
}

type linkBuf []rune

func (b linkBuf) IsEmpty() bool {
	return len(b) == 0
}

func (b linkBuf) IsEnd() bool {
	if b.IsEmpty() {
		return false
	}
	if len(b) >= 4 && b[len(b)-1] == ']' && b[len(b)-2] == ']' {
		return true
	}
	return false
}

func (b linkBuf) NewWikiFileLink() (*WikiFileLink, error) {
	l := new(WikiFileLink)
	s := strings.Trim(string(b), "[]")
	ss := strings.Split(s, "|")
	switch sl := len(ss); {
	case sl >= 3:
		l.Caption = ss[2]
		fallthrough
	case sl >= 2:
		l.Options = ss[1]
		fallthrough
	default:
		sss := strings.Split(ss[0], ":")
		if len(sss) > 1 {
			l.Type = sss[0]
			l.Name = sss[1]
		} else {
			l.Name = ss[0]
		}
	}
	return l, nil
}

func (s *WikiScanner) LookupLink() (*WikiFileLink, error) {
	r := s.r
	var buf linkBuf
	for {
		rr, _, err := r.ReadRune()
		if err != nil {
			return nil, err
		}
		switch rr {
		case '[':
			if buf.IsEmpty() || len(buf) == 1 {
				buf = append(buf, rr)
			} else {
				break
			}
		case ']':
			if len(buf) >= 2 && buf[len(buf)-3] != ']' {
				buf = append(buf, rr)
			} else {
				break
			}
		default:
			if len(buf) >= 2 && buf[len(buf)-1] != ']' {
				buf = append(buf, rr)
			}
		}
		if buf.IsEnd() {
			return buf.NewWikiFileLink()
		}
	}

	return nil, io.EOF
}
