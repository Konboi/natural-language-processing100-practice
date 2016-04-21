package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
cut -f1 hightemp.txt | LC_ALL=C sort | uniq -c | sort -r
*/

type Temp struct {
	Counter int
	Kenall  string
}

type TempSlice []Temp

func (t TempSlice) Len() int {
	return len(t)
}

func (t TempSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TempSlice) Less(i, j int) bool {
	return t[i].Counter > t[j].Counter
}

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("file path does not exist")
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

	m := make(map[string]int)
	n := make(map[string]int)
	for scanner.Scan() {
		lineSlice := strings.Split(scanner.Text(), "\t")
		if err != nil {
			fmt.Println(err.Error())
		}
		m[lineSlice[0]]++
		n[lineSlice[0]] = m[lineSlice[0]]
	}
	for k, v := range n {
		tempSlice = append(tempSlice, Temp{v, k})
	}
	sort.Sort(tempSlice)
	for _, temp := range tempSlice {
		fmt.Println(temp.Counter, temp.Kenall)
	}
}
