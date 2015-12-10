package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func ExampleNGram() {
	fmt.Println(NewFromNGram("paraparaparadise", 2))
	fmt.Println(NewFromNGram("paragraph", 2))
	fmt.Println(NewFromNGram2("日本語のテスト", 2))
	// Output:
	// (ad, ap, ar, di, is, pa, ra, se)
	// (ag, ap, ar, gr, pa, ph, ra)
	// (のテ, スト, テス, 日本, 本語, 語の)
}

func ExampleNGram2() {
	fmt.Println(NewFromNGram2("paraparaparadise", 2))
	fmt.Println(NewFromNGram2("paragraph", 2))
	fmt.Println(NewFromNGram2("日本語のテスト", 2))
	// Output:
	// (ad, ap, ar, di, is, pa, ra, se)
	// (ag, ap, ar, gr, pa, ph, ra)
	// (のテ, スト, テス, 日本, 本語, 語の)
}

// 和集合
func ExampleOr() {
	fmt.Println(NewFromNGram("paraparaparadise", 2).Or(NewFromNGram("paragraph", 2)))
	// Output:
	// (ad, ag, ap, ar, di, gr, is, pa, ph, ra, se)
}

// 積集合
func ExampleAnd() {
	fmt.Println(NewFromNGram("paraparaparadise", 2).And(NewFromNGram("paragraph", 2)))
	// Output:
	// (ap, ar, pa, ra)
}

// 差集合
func ExampleSub() {
	fmt.Println(NewFromNGram("paraparaparadise", 2).Sub(NewFromNGram("paragraph", 2)))
	// Output:
	// (ad, di, is, se)
}

// 差集合
func ExampleContains() {
	fmt.Println(NewFromNGram("paraparaparadise", 2).Contains("se"))
	fmt.Println(NewFromNGram("paragraph", 2).Contains("se"))
	// Output:
	// true
	// false
}

func ReadString(filename string) string {
	bytes, _ := ioutil.ReadFile(filename)
	return string(bytes)
}

func BenchmarkNGram(b *testing.B) {
	str := ReadString("wagahaiwa_nekodearu.txt")
	for i := 0; i < b.N; i++ {
		NewFromNGram(str, 2)
	}
}

func BenchmarkNGram2(b *testing.B) {
	str := ReadString("wagahaiwa_nekodearu.txt")
	for i := 0; i < b.N; i++ {
		NewFromNGram2(str, 2)
	}
}

func BenchmarkOr(b *testing.B) {
	s1 := NewFromNGram2(ReadString("wagahaiwa_nekodearu.txt"), 2)
	s2 := NewFromNGram2(ReadString("bocchan.txt"), 2)
	for i := 0; i < b.N; i++ {
		s1.Or(s2)
	}
}

func BenchmarkAnd(b *testing.B) {
	s1 := NewFromNGram2(ReadString("wagahaiwa_nekodearu.txt"), 2)
	s2 := NewFromNGram2(ReadString("bocchan.txt"), 2)
	for i := 0; i < b.N; i++ {
		s1.And(s2)
	}
}

func BenchmarkSub(b *testing.B) {
	s1 := NewFromNGram2(ReadString("wagahaiwa_nekodearu.txt"), 2)
	s2 := NewFromNGram2(ReadString("bocchan.txt"), 2)
	for i := 0; i < b.N; i++ {
		s1.Sub(s2)
	}
}

func BenchmarkContains(b *testing.B) {
	s := NewFromNGram2(ReadString("wagahaiwa_nekodearu.txt"), 2)
	for i := 0; i < b.N; i++ {
		s.Contains("ああ")
	}
}
