package main

//
// $ go build -o split_file
// $ ./split_file -n 3 hightemp.txt

// $ cat hightemp.txt | wc -l | xargs -J% split -l % test.txt _split.
//

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n = 2
)

func main() {
	flag.IntVar(&n, "n", 2, "split number")
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

	line := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	i := len(line) % n // 端数行
	j := len(line) - i
	k := j / n // 1ファイルあたりの行数

	//	count := 1
	//	number := 1
	data := []string{}
	number := 0
	for l := 0; l < len(line); l++ {
		data = append(data, line[l])
		if len(data) == k {
			if i > 0 {
				i--
				l++
				data = append(data, line[l])
			}
			number++
			filePath := fmt.Sprintf("./split.%d.txt", number)
			saveFile(filePath, data)
			data = []string{}
		}
	}
}

func saveFile(filePath string, data []string) {
	fout, _ := os.Create(filePath)
	defer fout.Close()
	for _, d := range data {
		fout.WriteString(fmt.Sprintf("%s\n", d))
	}
	fmt.Printf("> %s\n", filePath)
}
