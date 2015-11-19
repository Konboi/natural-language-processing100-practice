package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/*
14. 先頭からN行を出力

自然数Nをコマンドライン引数などの手段で受け取り，
入力のうち先頭のN行だけを表示せよ．
確認にはheadコマンドを用いよ．

*/

/*
$ head hightemp.txt
高知県  江川崎  41      2013-08-12
埼玉県  熊谷    40.9    2007-08-16
岐阜県  多治見  40.9    2007-08-16
山形県  山形    40.8    1933-07-25
山梨県  甲府    40.7    2013-08-10
和歌山県        かつらぎ        40.6    1994-08-08
静岡県  天竜    40.6    1994-08-04
山梨県  勝沼    40.5    2013-08-10
埼玉県  越谷    40.4    2007-08-16
群馬県  館林    40.3    2007-08-16

$ head -n 3 hightemp.txt
高知県  江川崎  41      2013-08-12
埼玉県  熊谷    40.9    2007-08-16
岐阜県  多治見  40.9    2007-08-16
*/

var (
	n = 0
)

func main() {
	flag.IntVar(&n, "n", 10, "lines")
	flag.Parse()

	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("Please Set File Name")
		return
	}

	filePath := fmt.Sprintf("./%s", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++

		if count >= n {
			return
		}
	}
}
