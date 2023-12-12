package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	printNumber bool
	printStat   bool
)

var (
	statTotalWords    int
	statTotalNewWords int
)

func init() {
	flag.BoolVar(&printNumber, "n", false, "print ordered number for each words")
	flag.BoolVar(&printStat, "s", false, "print statistics of words")
	flag.Parse()
}

func main() {
	files := flag.Args()

	myDict, err := OpenMyDict(myDictName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read my dict: %s", err.Error())
		os.Exit(1)
	}

	dict, err := NewWnDict("dict")
	if err != nil {
		fmt.Fprintf(os.Stderr, "new wn dict: %s", err.Error())
		os.Exit(1)
	}

	newWordsMap := make(map[string]struct{}, 10)
	newWords := make([]string, 0, 10)
	for _, fname := range files {
		text, err := os.ReadFile(fname)
		if err != nil {
			fmt.Printf("read file %s: %s\n", fname, err.Error())
			continue
		}

		tokenizer, err := NewTokenizer(string(text))
		if err != nil {
			fmt.Printf("new tokenizer %s: %s\n", fname, err.Error())
			continue
		}

		words := tokenizer.Tokens()
		for _, word := range words {
			word, err := dict.Morph(word)
			if err != nil {
				continue
			}
			known, err := myDict.Lookup(strings.ToLower(word))
			if err != nil {
				fmt.Fprintf(os.Stderr, "lookup mydict %s: %s", word, err.Error())
				continue
			}
			if known {
				continue
			}
			_, exists := newWordsMap[string(word)]
			if !exists {
				newWordsMap[string(word)] = struct{}{}
				newWords = append(newWords, word)
			}
		}
		statTotalWords += len(words)
	}
	statTotalNewWords = len(newWords)
	if printStat {
		fmt.Printf("total words: %d\n", statTotalWords)
		fmt.Printf("total new words: %d\n", statTotalNewWords)
	}

	if printNumber {
		for i, w := range newWords {
			fmt.Printf("%d. %s\n", i+1, w)
		}
	} else {
		for _, w := range newWords {
			fmt.Printf("%s\n", w)
		}
	}

}
