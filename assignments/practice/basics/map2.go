package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ss := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range ss {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	return m
}

func main() {
	// color := "yellow is good color"
	wc.Test(WordCount)
}
