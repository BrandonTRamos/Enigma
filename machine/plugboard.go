package machine

import (
	"strings"
)

type PlugBoard struct {
	mappings map[rune]rune
}

func NewPlugBoard() *PlugBoard {
	plugboard := PlugBoard{}
	plugboard.mappings = make(map[rune]rune)
	return &plugboard
}

func NewPlugBoardFromPairString(pairs string) *PlugBoard {
	plugboard := NewPlugBoard()
	if pairs != "" {
		split := strings.Split(pairs, ",")
		for _, pair := range split {
			plugboard.addMappingPair(rune(pair[0]), rune(pair[1]))
		}
	}
	return plugboard
}

func (p *PlugBoard) addMappingPair(firstLetter rune, secondLetter rune) {
	p.mappings[firstLetter] = secondLetter
	p.mappings[secondLetter] = firstLetter
}
func (p *PlugBoard) removeMappingPair(firstLetter rune, secondLetter rune) {
	delete(p.mappings, firstLetter)
	delete(p.mappings, secondLetter)
}

func (p *PlugBoard) SwapLetter(letter rune) rune {
	swap, present := p.mappings[letter]
	if present {
		return swap
	}
	return letter
}
