package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	COL1_FILENAME   = "col1.txt"
	COL2_FILENAME   = "col2.txt"
	MERGED_FILENAME = "merged.txt"
)

func main() {
	MergeColumns()
}

func MergeColumnsWithChannel() {
	ch1 := make(chan string, 256)
	ch2 := make(chan string, 256)

	f1, err := os.Open(COL1_FILENAME)
	if err != nil {
		log.Fatalln("open col1.txt error:", err)
	}
	defer f1.Close()
	f2, err := os.Open(COL2_FILENAME)
	if err != nil {
		log.Fatalln("open col2.txt error:", err)
	}
	defer f2.Close()

	mf, err := os.Create(MERGED_FILENAME)
	if err != nil {
		log.Fatalln("create merged.txt error:", err)
	}
	defer mf.Close()

	go ScanFile(f1, ch1)
	go ScanFile(f2, ch2)
	for {
		c1, ok1 := <-ch1
		if !ok1 {
			break
		}
		c2, ok2 := <-ch2
		if !ok2 {
			break
		}
		fmt.Fprintf(mf, "%s\t%s\n", c1, c2)
	}
}

func ScanFile(r io.Reader, ch chan string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		ch <- t
	}
	close(ch)
}

func MergeColumns() {
	f1, err := os.Open(COL1_FILENAME)
	if err != nil {
		log.Fatalln("open col1.txt error:", err)
	}
	defer f1.Close()
	s1 := bufio.NewScanner(f1)
	f2, err := os.Open(COL2_FILENAME)
	if err != nil {
		log.Fatalln("open col2.txt error:", err)
	}
	defer f2.Close()
	s2 := bufio.NewScanner(f2)

	mf, err := os.Create(MERGED_FILENAME)
	if err != nil {
		log.Fatalln("create merged.txt error:", err)
	}
	defer mf.Close()

	for s1.Scan() && s2.Scan() {
		c1 := s1.Text()
		c2 := s2.Text()
		fmt.Fprintf(mf, "%s\t%s\n", c1, c2)
	}
}
