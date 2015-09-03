package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// hightemp.txtは，日本の最高気温の記録を「都道府県」「地点」「℃」「日」のタブ区切り形式で格納したファイルである．
// 以下の処理を行うプログラムを作成し，hightemp.txtを入力ファイルとして実行せよ．
// さらに，同様の処理をUNIXコマンドでも実行し，プログラムの実行結果を確認せよ．

// unix command
// cat higtemp.txt | wc -l
// 24

// go run main.go
// 24

var (
	count = 0
)

func main() {
	fileName := "hightemp.txt"
	filePath := fmt.Sprintf("./%s", fileName)

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		_, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		count++
	}

	fmt.Printf("%d \n", count)
}
