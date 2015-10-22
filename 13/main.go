package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	COL1_FILENAME   = "col1.txt"
	COL2_FILENAME   = "col2.txt"
	MERGED_FILENAME = "merged.txt"
)

var (
	ch1 = make(chan string)
	ch2 = make(chan string)
	wg  sync.WaitGroup
)

func main() {
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

	wg.Add(2)
	go ScanFile(f1, ch1)
	go ScanFile(f2, ch2)
	go func() {
		for {
			var c1, c2 string
			for c1 == "" || c2 == "" {
				select {
				case t := <-ch1:
					c1 = t
				case t := <-ch2:
					c2 = t
				}
			}
			fmt.Fprintf(mf, "%s\t%s\n", c1, c2)
		}
	}()
	wg.Wait()
}

func ScanFile(r io.Reader, ch chan string) {
	defer wg.Done()
	s := bufio.NewScanner(r)
	for s.Scan() {
		t := s.Text()
		ch <- t
	}
}
