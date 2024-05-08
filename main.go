package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	printNumber  bool
	printStat    bool
	showPercents int
)

var (
	statTotalWords    int
	statTotalNewWords int
)

func init() {
	flag.BoolVar(&printNumber, "n", false, "print ordered number for each words")
	flag.BoolVar(&printStat, "s", false, "print statistics of words")
	flag.IntVar(&showPercents, "p", 20, "show percents")
	flag.Parse()
}

func main() {
	files := flag.Args()
	myDict, err := OpenMyDict(myDictName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read my dict: %s", err.Error())
		os.Exit(1)
	}

	dict, err := NewTxtDict("dict")
	if err != nil {
		fmt.Fprintf(os.Stderr, "new wn dict: %s", err.Error())
		os.Exit(1)
	}

	minShowRank := (showPercents * dict.Count()) / 100.0

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
			idx, wd := dict.Lookup(strings.ToLower(word))
			if idx < minShowRank {
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
			_, exists := newWordsMap[wd.lemma]
			if !exists {
				newWordsMap[wd.lemma] = struct{}{}
				newWords = append(newWords, wd.lemma)
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
