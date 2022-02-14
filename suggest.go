package main

import (
	"fmt"
	"log"
	"regexp"
)

func containBites(word string, bites string) bool {
candidate:
	for _, candidate := range bites {
		for _, c := range word {
			if c == candidate {
				continue candidate
			}
		}
		return false
	}
	return true
}

func suggestWords(answers []string) []string {
	bites := ""
	notUsed := ""
	rss := [5]string{}
	for _, v := range answers {
		i := 0
		for j := 0; j < len(v); j++ {
			if v[j] == '!' {
				j++
				rss[i] = string(v[j])
			} else if v[j] == '?' {
				j++
				if rss[i] == "*" || rss[i] == "" {
					rss[i] = fmt.Sprintf("?%c", v[j])
				} else {
					rss[i] += string(v[j])
				}
				bites += string(v[j])
			} else {
				if rss[i] == "" {
					rss[i] += "*"
				}
				notUsed += string(v[j])
			}
			i++
		}
	}

	for i := 0; i < len(rss); i++ {
		if rss[i] == "*" {
			rss[i] = fmt.Sprintf("[^%s]", notUsed)
		} else if rss[i] == "?" {
			rss[i] = fmt.Sprintf("[^%s%s]", notUsed, rss[i][1:])
		}
	}

	reg, err := regexp.Compile(`\?(.*)`)
	if err != nil {
		log.Fatal(err)
	}
	rs := ""
	for i := range rss {
		rss[i] = reg.ReplaceAllString(rss[i], fmt.Sprintf("[^%s$1]", notUsed))
		rs += rss[i]
	}

	words := getWords()

	var res []string
	reg, err = regexp.Compile(rs)
	for _, v := range words {
		if reg.MatchString(v) && containBites(v, bites) {
			res = append(res, v)
		}
	}
	return res
}
