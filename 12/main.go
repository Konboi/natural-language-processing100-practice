package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 各行の1列目だけを抜き出したものをcol1.txtに，
// 2列目だけを抜き出したものをcol2.txtとしてファイルに保存せよ．
// 確認にはcutコマンドを用いよ．

// cut -f 1 hightemp.txt
// cut -f 2 hightemp.txt

var (
	col1 = ""
	col2 = ""
)

func main() {
	fileName := "hightemp.txt"
	filePath := fmt.Sprintf("./%s", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columns := strings.Split(scanner.Text(), "\t")
		col1 = col1 + columns[0] + "\n"
		col2 = col2 + columns[1] + "\n"
	}
	ioutil.WriteFile("./col1.txt", []byte(col1), os.ModePerm)
	ioutil.WriteFile("./col2.txt", []byte(col2), os.ModePerm)
}
