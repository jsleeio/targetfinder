package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func match(word, target, key string) bool {
	// words must be 4-9 letters
	if len(word) < 4 || len(word) > 9 {
		return false
	}
	// words must contain the key letter
	if !strings.Contains(word, key) {
		return false
	}
	depletingTarget := target
	runeTest := func(r rune) rune {
		if strings.ContainsRune(depletingTarget, r) {
			// each letter may be used only once
			depletingTarget = strings.Replace(depletingTarget, string(r), "", 1)
			return r
		}
		return '!'
	}
	return strings.Map(runeTest, word) == word
}

func dictSubset(dict, target, key string) []string {
	words := []string{}
	fd, err := os.Open(dict)
	defer fd.Close()
	if err != nil {
		fmt.Printf("error reading dictionary: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := scanner.Text()
		if match(word, target, key) {
			words = append(words, word)
		}
	}
	return words
}

func main() {
	var dict, target, key string
	flag.StringVar(&dict, "dictionary", "/usr/share/dict/words", "location of system dictionary file")
	flag.StringVar(&target, "target", "dtgmtiaei", "target letters. Must be exactly 9 characters")
	flag.StringVar(&key, "key", "t", "key letter. Must be exactly 1 character")
	flag.Parse()
	if len(target) != 9 {
		fmt.Printf("target must be exactly 9 letters, got: %q\n", target)
		os.Exit(1)
	}
	if len(key) != 1 {
		fmt.Printf("key must be exactly 1 letter, got: %q\n", key)
		os.Exit(1)
	}
	for _, word := range dictSubset(dict, target, key) {
		if len(word) == 9 {
			fmt.Println(strings.ToUpper(word))
		} else {
			fmt.Println(word)
		}
	}
}
