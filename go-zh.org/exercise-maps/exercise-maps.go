package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	r := strings.Fields(s)
	for _, v := range r {
		ele, ok := result[v]
		if ok {
			result[v] = ele + 1
		} else {
			result[v] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
