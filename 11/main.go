package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//sed -e 's/\t/ /g' ./hightemp.txt

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
		text := scanner.Text()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		resStr := strings.Replace(text, "\t", " ", -1)
		fmt.Println(resStr)
	}
}
