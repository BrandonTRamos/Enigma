package machine


type Rotor struct {
	name                string
	encoding            string
	forwardWiring       [26]int32
	reverseWiring       [26]int32
	rotorPositionOffset int32
	ringSettingOffset   int32
}

func NewRotor(name string, encoding string,ringSettingOffset int32) *Rotor {
	rotor := Rotor{name: name, encoding: encoding,ringSettingOffset: ringSettingOffset}
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
	r.rotorPositionOffset = (26 + (r.rotorPositionOffset+1)) % 26
}

func (r *Rotor) encodeForward(input int32) int32 {
	//return r.forwardWiring[input]
	ringOffSet:=((26+r.ringSettingOffset))%26
	return r.forwardWiring[(input+ringOffSet)%26]
}

func (r *Rotor) encodeReverse(input int32) int32 {
	ringOffSet:=((26+r.ringSettingOffset))%26
	return (26+(r.reverseWiring[input])-ringOffSet)%26
}

func (r *Rotor) GetCurrentRotorOffset() int32 {
	return r.rotorPositionOffset
}
