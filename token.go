package main

import (
	"regexp"

	"github.com/jdkato/prose/v3"
)

var (
	suffixes    = []string{",", ")", `"`, "]", "!", ";", ".", "?", ":", "'", "*"}
	prefixes    = []string{"$", "(", `"`, "[", "-", "*"}
	reValidWord = regexp.MustCompile(`^[a-zA-Z]+[-']*[a-zA-Z]*$`)
)

type Tokenizer struct {
	doc *prose.Document
}

func NewTokenizer(text string) (*Tokenizer, error) {
	tokenizer := prose.NewIterTokenizer(prose.UsingPrefixes(prefixes), prose.UsingSuffixes(suffixes))
	doc, err := prose.NewDocument(string(text), prose.UsingTokenizer(tokenizer))
	if err != nil {
		return nil, err
	}

	return &Tokenizer{doc: doc}, nil
}

func (to *Tokenizer) Tokens() []string {
	var ts = make([]string, 0, 1024)
	for _, t := range to.doc.Tokens() {
		if to.isValidWord(t) {
			ts = append(ts, t.Text)
		}
	}

	return ts
}

func (to *Tokenizer) isValidWord(t prose.Token) bool {
	return reValidWord.Match([]byte(t.Text))
}
