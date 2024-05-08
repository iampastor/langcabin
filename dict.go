package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

type wd struct {
	morph string
	lemma string
}

type TxtDict struct {
	words []wd
}

func NewTxtDict(dir string) (*TxtDict, error) {
	file, err := os.Open(filepath.Join(dir, "eng_dict.data"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = '\t'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	words := make([]wd, 0, len(records))
	for _, row := range records {
		word := wd{
			morph: row[0],
			lemma: row[1],
		}
		words = append(words, word)
	}

	return &TxtDict{
		words: words,
	}, nil
}

func (d *TxtDict) Lookup(word string) (int, wd) {
	for i, w := range d.words {
		if w.morph == word {
			return i, w
		}
	}

	return -1, wd{}
}

func (d *TxtDict) Count() int {
	return len(d.words)
}
