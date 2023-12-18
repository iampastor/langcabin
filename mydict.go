package main

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
)

var (
	myDictName = "mydict"
)

// MyDict contains words that we alreay know
type MyDict struct {
	dict map[string]bool
}

func OpenMyDict(dirname string) (*MyDict, error) {
	fnames, err := filepath.Glob(filepath.Join(dirname, "*.data"))
	if err != nil {
		return nil, err
	}
	dict := make(map[string]bool, 5000)
	for _, fname := range fnames {
		dictData, err := os.ReadFile(fname)
		if err != nil {
			return nil, err
		}
		dictWords := bytes.Split(dictData, []byte("\n"))
		for _, w := range dictWords {
			dict[string(w)] = true
		}
	}

	return &MyDict{dict: dict}, nil
}

func (d *MyDict) Lookup(word string) (bool, error) {
	return d.dict[word], nil
}

func (d *MyDict) AddWord(word string) error {
	d.dict[word] = true
	return nil
}

func (d *MyDict) RemoveWord(word string) error {
	delete(d.dict, word)
	return nil
}

func (d *MyDict) Save() error {
	data := make([]string, 0, len(d.dict))
	for w := range d.dict {
		data = append(data, w)
	}
	sort.Strings(data)

	f, err := os.OpenFile("dict.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	for _, w := range data {
		_, err = f.Write([]byte(w + "\n"))
		if err != nil {
			return err
		}
	}
	f.Close()
	return os.Rename("dict.tmp", "dict.data")
}
