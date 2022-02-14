package main

import (
	"fmt"
	"os"
	"wordle-tool"
)

func main() {
	var answers []string
	for i := 0; i < len(os.Args)-1; i++ {
		answers = append(answers, os.Args[i+1])
	}

	res := wordle_cli.SuggestWords(answers)

	fmt.Println("\nsuggest:")
	for i, v := range res {
		fmt.Printf("%s ", v)
		if i != 0 && i%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
