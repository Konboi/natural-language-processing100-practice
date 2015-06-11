package main

/*
BenchmarkZipStr            2000000               694 ns/op
BenchmarkZipStrGoroutine    100000             14748 ns/op
BenchmarkZipStrArr         2000000               836 ns/op
BenchmarkZipStrOneLoop     2000000               842 ns/op
*/

import (
	"testing"
)

var (
	testWord1 = "パトカー"
	testWord2 = "タクシー"
)

func TestZipStr(t *testing.T) {
	result := ZipStr(testWord1, testWord2)
	if result != "パタトクカシーー" {
		t.Error("not match パタトクカシーー:", result)
	}
}

func BenchmarkZipStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZipStr(testWord1, testWord2)
	}
}

func TestZipStrByGoroutine(t *testing.T) {
	result := ZipStrByGoroutine(testWord1, testWord2)
	if result != "パタトクカシーー" {
		t.Error("not match パタトクカシーー:", result)
	}
}

func BenchmarkZipStrByGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZipStrByGoroutine(testWord1, testWord2)
	}
}

func TestZipStrArr(t *testing.T) {
	result := ZipStrArr(testWord1, testWord2)
	if result != "パタトクカシーー" {
		t.Error("not match パタトクカシーー:", result)
	}
}

func BenchmarkZipStrArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZipStrArr(testWord1, testWord2)
	}
}

func TestZipStrOneLoop(t *testing.T) {
	result := ZipStrOneLoop(testWord1, testWord2)
	if result != "パタトクカシーー" {
		t.Error("not match パタトクカシーー:", result)
	}
}

func BenchmarkZipStrOneLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZipStrOneLoop(testWord1, testWord2)
	}
}
