package main

import (
	"fmt"
	"os"
)

func main() {
	var answers []string
	for i := 0; i < len(os.Args)-1; i++ {
		answers = append(answers, os.Args[i+1])
	}

	res := suggestWords(answers)

	fmt.Println("suggest:")
	for i, v := range res {
		fmt.Printf("%s ", v)
		if i != 0 && i%4 == 0 {
			fmt.Println()
		}
	}
}
