package main

/*
enchmarkCountWordStrings         1000000              1565 ns/op
BenchmarkCountWordStringsTrim     500000              3321 ns/op
BenchmarkCountWordStringsMattn   1000000              1511 ns/op
BenchmarkCountWordStringsMethane 2000000               989 ns/op
*/

import (
	"reflect"
	"testing"
)

var expectedCounts = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}
var testSentence = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

func TestCountWordStrings(t *testing.T) {
	counts := CountWordStrings(testSentence)
	if !reflect.DeepEqual(counts, expectedCounts) {
		t.Error("this is no pi:", counts)
	}
}

func BenchmarkCountWordStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWordStrings(testSentence)
	}
}

func TestCountWordStringsTrim(t *testing.T) {
	counts := CountWordStringsTrim(testSentence)
	if !reflect.DeepEqual(counts, expectedCounts) {
		t.Error("this is no pi:", counts)
	}
}

func BenchmarkCountWordStringsTrim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWordStringsTrim(testSentence)
	}
}

func TestCountWordStringsMattn(t *testing.T) {
	counts := CountWordStringsMattn(testSentence)
	if !reflect.DeepEqual(counts, expectedCounts) {
		t.Error("this is no pi:", counts)
	}
}

func BenchmarkCountWordStringsMattn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWordStringsMattn(testSentence)
	}
}

func TestCountWordStringsMethane(t *testing.T) {
	counts := CountWordStringsMethane(testSentence)
	if !reflect.DeepEqual(counts, expectedCounts) {
		t.Error("this is no pi:", counts)
	}
}

func BenchmarkCountWordStringsMethane(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWordStringsMethane(testSentence)
	}
}
