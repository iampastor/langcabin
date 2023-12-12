package main

import "github.com/ozlo/gown"

type WnDict struct {
	wn *gown.WN
}

func NewWnDict(dictDir string) (*WnDict, error) {
	wn, err := gown.LoadWordNet(dictDir)
	if err != nil {
		return nil, err
	}
	wn.InitMorphData("dict")
	return &WnDict{wn: wn}, nil
}

func (d *WnDict) Morph(word string) (string, error) {
	if len(d.wn.Lookup(word)) != 0 {
		return word, nil
	}

	vw := d.wn.Morph(word, gown.POS_VERB)
	nw := d.wn.Morph(word, gown.POS_NOUN)
	if vw != "" {
		return vw, nil
	}
	if nw != "" {
		return nw, nil
	}

	return word, nil
}
