package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
各行を3コラム目の数値の逆順で整列せよ（注意: 各行の内容は変更せずに並び替えよ）
確認にはsortコマンドを用いよ（この問題はコマンドで実行した時の結果と合わなくてもよい）．
*/

/*
$ sort -n -k 3 -t $'\t' hightemp.txt
*/

type Temp struct {
	Pos1     string
	Pos2     string
	Hightemp float64
	Date     string
	Text     string
}

type TempSlice []Temp

// sort.Interfaceを満たすためLen(), Swap(), Less()を定義
func (t TempSlice) Len() int {
	return len(t)
}

func (t TempSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TempSlice) Less(i, j int) bool {
	return t[i].Hightemp < t[j].Hightemp
}

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("please set file name.")
		return
	}

	filePath := fmt.Sprintf("./%s", fileName)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	tempSlice := TempSlice{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSlice := strings.Split(scanner.Text(), "\t")
		hightemp, err := strconv.ParseFloat(lineSlice[2], 64)
		if err != nil {
			fmt.Println(err.Error())
		}

		tempSlice = append(tempSlice, Temp{lineSlice[0], lineSlice[1], hightemp, lineSlice[3], scanner.Text()})
	}

	sort.Sort(tempSlice)

	for _, temp := range tempSlice {
		fmt.Println(temp.Text)
	}
}
