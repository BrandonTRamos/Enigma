package machine

type Rotor struct {
	Name          string
	Encoding      string
	ForwardWiring [26]int32
	ReverseWiring [26]int32
	RotorPosition int32
	RingSetting   int32
	RotationPoint int32
}

var rotorRotationPointMap map[string]int32 = map[string]int32{"I": 16, "II": 4, "III": 21}
var rotorNumberEncodingMap map[string]string = map[string]string{"I": FirstRotorEncoding, "II": SecondRotorEncoding, "III": ThirdRotorEncoding}

func NewRotorFromName(name string) *Rotor {
	rotor := Rotor{Name: name, Encoding: rotorNumberEncodingMap[name], RotationPoint: rotorRotationPointMap[name]}
	rotor.generateForwardWiring()
	rotor.generateReverseWiring()
	return &rotor
}

func NewRotorFromNameAndRotorPosition(name string, rotorPosition int32) *Rotor {
	rotor := NewRotorFromName(name)
	rotor.RotorPosition = rotorPosition
	return rotor
}

func NewRotor(name string, encoding string, rotorPosition int32, ringSettingOffset int32, rotationPoint int32) *Rotor {
	rotor := Rotor{Name: name, Encoding: encoding, RingSetting: ringSettingOffset, RotorPosition: rotorPosition, RotationPoint: rotationPoint}
	rotor.generateForwardWiring()
	rotor.generateReverseWiring()
	return &rotor
}

func (r *Rotor) generateForwardWiring() {
	for idx, char := range r.Encoding {
		r.ForwardWiring[idx] = (char - 65)
	}
}
func (r *Rotor) generateReverseWiring() {
	for idx, char := range r.ForwardWiring {
		r.ReverseWiring[char] = int32(idx)
	}
}

func (r *Rotor) rotate() {
	r.RotorPosition = (26 + (r.RotorPosition + 1)) % 26
}

func (r *Rotor) encodeForward(input int32) int32 {

	offset := (26 + (r.RotorPosition - r.RingSetting)) % 26
	return r.ForwardWiring[(input+offset)%26]
}

func (r *Rotor) encodeReverse(input int32) int32 {
	offset := (26 + (r.RotorPosition - r.RingSetting)) % 26
	return (26 + (r.ReverseWiring[input]) - offset) % 26
}

func (r *Rotor) GetCurrentRotorOffset() int32 {
	return r.RotorPosition
}

func (r *Rotor) GetName() string {
	return r.Name
}

func (r *Rotor) SetRotorPosition(position int32) {
	r.RotorPosition = position
}
func (r *Rotor) GetRingSetting() int32 {
	return r.RingSetting
}
