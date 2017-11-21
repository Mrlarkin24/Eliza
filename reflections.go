package main

import (
	"fmt"
	"regexp"
	"strings"
)


func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)
	
	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}

func main() {
	fmt.Println(Reflect("You are my friend."))
}