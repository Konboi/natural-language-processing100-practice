package main

import (
	"testing"
)

func BenchmarkMergeColumns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeColumns()
	}
}

func BenchmarkMergeColumnsWithChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeColumnsWithChannel()
	}
}
