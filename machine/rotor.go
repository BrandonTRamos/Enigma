package machine

type Rotor struct {
	name               string
	encoding           string
	forwardWiring      [26]int32
	reverseWiring      [26]int32
	ringPositionOffset int32
}

func NewRotor(name string, encoding string) *Rotor {
	rotor := Rotor{name: name, encoding: encoding}
	rotor.generateForwardWiring()
	rotor.generateReverseWiring()
	return &rotor
}

func (r *Rotor) generateForwardWiring() {
	for idx, char := range r.encoding {
		r.forwardWiring[idx] = (char - 65)
	}
}
func (r *Rotor) generateReverseWiring() {
	for idx, char := range r.forwardWiring {
		r.reverseWiring[char] = int32(idx)
	}
}

func (r *Rotor) rotate() {
	r.ringPositionOffset = (26 + r.ringPositionOffset) % 26
}

func (r *Rotor) encodeForward(input int32) int32 {
	return r.forwardWiring[input]
}

func (r *Rotor) encodeReverse(input int32) int32 {
	return r.reverseWiring[input]
}
