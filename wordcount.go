package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {

	fmt.Println("initial string:", s)
	strings.TrimSpace(s)
	items := strings.Fields(s)
	wordmap := make(map[string]int)
	for _, j := range items {

		if wordmap[j] != 0 {

			wordmap[j]++

		} else {

			wordmap[j] = 1

		}

	}

	for i, j := range wordmap {

		fmt.Printf("%s-------%d\n", i, j)

	}

	return wordmap
}

func main() {

	wordCount(" shaofeng zhang is a very good guy, he is my classmates, his wife si helen zhang also very good,sister is in the same school with his daughter")

}
