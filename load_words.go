package wordle_cli

import (
	"encoding/json"
	"log"
	"os"
)

type Words struct {
	W []string `json:"words"`
}

func LoadWords(sp *[]string, fn string) {
	b, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var w Words
	if err = json.Unmarshal(b, &w); err != nil {
		log.Fatal(err)
	}
	*sp = w.W
}
