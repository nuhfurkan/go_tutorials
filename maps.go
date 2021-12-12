package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	newField := strings.Fields(s)
	myMap := make(map[string]int)
	for _, value := range newField {
		_, ok := myMap[value]
		if (ok == true) {
			myMap[value] += 1
		} else {
			myMap[value] = 1
		}
	}
	return myMap
}

func main() {
	wc.Test(WordCount)
}

