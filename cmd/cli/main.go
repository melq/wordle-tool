package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"wordle-tool"
)

func containCandidates(word string, candidates string) bool {
candidate:
	for _, candidate := range candidates {
		for _, c := range word {
			if c == candidate {
				continue candidate
			}
		}
		return false
	}
	return true
}

func main() {
	reg, err := regexp.Compile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	candidates := os.Args[2]

	var words []string
	wordle_cli.LoadWords(&words)

	var res []string
	for _, v := range words {
		if reg.MatchString(v) && containCandidates(v, candidates) {
			res = append(res, v)
		}
	}

	fmt.Println("suggest:")
	for i, v := range res {
		fmt.Printf("%s ", v)
		if i != 0 && i%4 == 0 {
			fmt.Println()
		}
	}
}
