package machine


type Reflector struct {
	encoding string
	wiring   [26]int32
}

func (r *Reflector) Reflect(input int32) int32 {
	//fmt.Println("Reflect: ", input,"->",r.wiring[input])
	return r.wiring[input]
}

func (r *Reflector) generateWiring() {
	for idx, char := range r.encoding {
		r.wiring[idx] = (char - 65)
	}
}

func NewReflector() *Reflector {
	encoding := "ZYXWVUTSRQPONMLKJIHGFEDCBA"
	reflector := &Reflector{encoding: encoding}
	reflector.generateWiring()
	return reflector
}
