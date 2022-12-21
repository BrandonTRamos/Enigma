package machine

type Rotor struct {
	name          string
	encoding      string
	forwardWiring [26]int32
	reverseWiring [26]int32
	rotorPosition int32
	ringSetting   int32
	rotationPoint int32
}

var rotorRotationPointMap map[string]int32 = map[string]int32{"I": 16, "II": 4, "III": 21}
var rotorNumberEncodingMap map[string]string = map[string]string{"I": firstRotorEncoding, "II": secondRotorEncoding, "III": thirdRotorEncoding}

func NewRotorFromName(name string) *Rotor {
	rotor := Rotor{name: name, encoding: rotorNumberEncodingMap[name], rotationPoint: rotorRotationPointMap[name]}
	rotor.generateForwardWiring()
	rotor.generateReverseWiring()
	return &rotor
}

func NewRotor(name string, encoding string, rotorPositionOffset int32, ringSettingOffset int32, rotationPoint int32) *Rotor {
	rotor := Rotor{name: name, encoding: encoding, ringSetting: ringSettingOffset, rotorPosition: rotorPositionOffset, rotationPoint: rotationPoint}
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
	r.rotorPosition = (26 + (r.rotorPosition + 1)) % 26
}

func (r *Rotor) encodeForward(input int32) int32 {

	offset := (26 + (r.rotorPosition - r.ringSetting)) % 26
	return r.forwardWiring[(input+offset)%26]
}

func (r *Rotor) encodeReverse(input int32) int32 {
	offset := (26 + (r.rotorPosition - r.ringSetting)) % 26
	return (26 + (r.reverseWiring[input]) - offset) % 26
}

func (r *Rotor) GetCurrentRotorOffset() int32 {
	return r.rotorPosition
}

func (r *Rotor) GetName() string {
	return r.name
}
