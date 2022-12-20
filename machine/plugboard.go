package machine

type PlugBoard struct {
	mappings map[rune]rune
}

func NewPlugBoard() *PlugBoard {
	plugboard := PlugBoard{}
	plugboard.mappings = make(map[rune]rune)
	return &plugboard
}

func (p *PlugBoard) addMappingPair(firstLetter rune, secondLetter rune) {
	p.mappings[firstLetter] = secondLetter
	p.mappings[secondLetter] = firstLetter
}

func (p *PlugBoard) SwapLetter(letter rune) rune {
	swap, present := p.mappings[letter]
	if present {
		return swap
	}
	return letter
}
