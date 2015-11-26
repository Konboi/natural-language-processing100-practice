package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/*
15. 末尾のN行を出力
自然数Nをコマンドライン引数などの手段で受け取り，入力のうち末尾のN行だけを表示せよ．確認にはtailコマンドを用いよ．
*/

func main() {
	var num int64
	flag.Int64Var(&num, "n", 10, "number of lines")
	flag.Parse()

	// open file
	var reader io.Reader
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	// read lines
	buf := make([]string, num)
	scanner := bufio.NewScanner(reader)
	var lineno int64 = 0
	for scanner.Scan() {
		buf[lineno%num] = scanner.Text()
		lineno++
	}

	// output
	if lineno <= num {
		for _, line := range buf[0:lineno] {
			fmt.Println(line)
		}
	} else {
		for i := lineno; i < lineno+num; i++ {
			fmt.Println(buf[i%num])
		}
	}
}
