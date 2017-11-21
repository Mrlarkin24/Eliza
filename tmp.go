package main

import (
  "fmt"
  "regexp"
  "os"
  "log"
  "bufio"
  "strings"
)


type replacer struct {
	original    *regexp.Regexp
	replacement []string
}

func readreplacers(path string) []replacer {
	file, err := os.Open(path)
  if err != nil {
		log.Fatal(err)
	}
  defer file.Close()
  
  replacers := []replacer{}

  scanner := bufio.NewScanner(file)
  for newsection := true; scanner.Scan(); {
		switch {
		case strings.HasPrefix(scanner.Text(), "#"):
		case len(scanner.Text()) == 0:
			newsection = true
		case newsection == true:
			replacers = append(replacers, replacer{original: regexp.MustCompile(scanner.Text())})
			newsection = false
		default:
			replacers[len(replacers)-1].replacement = append(replacers[len(replacers)-1].replacement, scanner.Text())
		}
	}
   
  return replacers
}

func main() {
  responses := readreplacers("responses.txt")
  substitutions := readreplacers("responses.txt")
  
  fmt.Println(responses)
  fmt.Println(substitutions)
}