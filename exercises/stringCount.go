package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	fields := strings.Fields(s)
	mp := make(map[string]int)
	for _, v := range fields {
		elem, present := mp[v]
		if present {
			elem++
			mp[v] = elem
		} else {
			mp[v] = 1

		}

	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
