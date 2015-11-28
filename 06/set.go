package main

import (
	"sort"
	"strings"
)

type Set map[string]struct{}

func NewFromNGram(str string, n int) Set {
	runes := []rune(str)
	set := Set{}
	for i, j := 0, n; j <= len(runes); i, j = i+1, j+1 {
		set[string(runes[i:j])] = struct{}{}
	}
	return set
}

func NewFromNGram2(str string, n int) Set {
	set := Set{}
	posList := make([]int, n)
	i := 0
	for pos := range str {
		j := i % n
		if i >= n {
			set[str[posList[j]:pos]] = struct{}{}
		}
		posList[j] = pos
		i++
	}
	set[str[posList[i%n]:len(str)]] = struct{}{}
	return set
}

func (a Set) Or(b Set) Set {
	var size int
	if len(a) > len(b) {
		size = len(a)
	} else {
		size = len(b)
	}
	c := make(Set, size)
	for s := range a {
		c[s] = struct{}{}
	}
	for s := range b {
		c[s] = struct{}{}
	}
	return c
}

func (a Set) And(b Set) Set {
	c := Set{}
	if len(a) > len(b) {
		for s := range b {
			if _, ok := a[s]; ok {
				c[s] = struct{}{}
			}
		}
	} else {
		for s := range a {
			if _, ok := b[s]; ok {
				c[s] = struct{}{}
			}
		}
	}
	return c
}

func (a Set) Sub(b Set) Set {
	var size int
	if len(a) > len(b) {
		size = len(a) - len(b)
	} else {
		size = 0
	}
	c := make(Set, size)
	for s := range a {
		if _, ok := b[s]; !ok {
			c[s] = struct{}{}
		}
	}
	return c
}

func (a Set) Contains(str string) bool {
	_, ok := a[str]
	return ok
}

func (a Set) String() string {
	elems := make([]string, 0, len(a))
	for elem := range a {
		elems = append(elems, elem)
	}
	sort.Stable(sort.StringSlice(elems))
	return "(" + strings.Join(elems, ", ") + ")"
}
