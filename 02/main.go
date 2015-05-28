package main

import (
	"fmt"
	"unicode/utf8"
)

/*
02. 「パトカー」＋「タクシー」＝「パタトクカシーー」
「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ．
*/

var (
	patcar = "パトカー"
	taxi   = "タクシー"
)

func main() {
	zipped := ZipStr(patcar, taxi)
	fmt.Println(zipped)
}

func ZipStr(s1, s2 string) string {
	zippedLength := utf8.RuneCountInString(s1) + utf8.RuneCountInString(s2)
	zipped := make([]rune, zippedLength)
	i := 0
	for _, r := range s1 {
		zipped[i*2] = r
		i++
	}
	j := 0
	for _, r := range s2 {
		zipped[j*2+1] = r
		j++
	}
	return string(zipped)
}

func ZipStrArr(s1, s2 string) string {
	rs1 := []rune(s1)
	rs2 := []rune(s2)
	zippedLength := len(rs2) + len(rs1)
	zipped := make([]rune, zippedLength)
	for i, r := range rs1 {
		zipped[i*2] = r
	}
	for i, r := range rs2 {
		zipped[i*2+1] = r
	}
	return string(zipped)
}

func ZipStrOneLoop(s1, s2 string) string {
	rs1 := []rune(s1)
	rs2 := []rune(s2)
	zippedLength := len(rs2) + len(rs1)
	zipped := make([]rune, zippedLength)
	for i := 0; i < len(rs1); i++ {
		j := i * 2
		zipped[j], zipped[j+1] = rs1[i], rs2[i]
	}
	return string(zipped)
}

func ZipStrByGoroutine(s1, s2 string) string {
	zippedLength := utf8.RuneCountInString(s1) + utf8.RuneCountInString(s2)
	zipped := make([]rune, zippedLength)
	recvCh := make(chan rune)
	switchCh1 := make(chan struct{})
	switchCh2 := make(chan struct{})

	go func() {
		for _, r := range s1 {
			<-switchCh1
			recvCh <- r
			switchCh2 <- struct{}{}
		}
	}()
	go func() {
		for _, r := range s2 {
			<-switchCh2
			recvCh <- r
			switchCh1 <- struct{}{}
		}
	}()

	i := 0
	switchCh1 <- struct{}{}
	for r := range recvCh {
		zipped[i] = r
		i++
		if i == zippedLength {
			break
		}
	}

	return string(zipped)
}
