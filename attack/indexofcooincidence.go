package attack

// english text IOC ~ 0.0686
func CalcIndexOfCooincidence(text string) float32 {
	var histogram [26]int32
	for _, char := range text {
		histogram[char-65]++
	}
	length := float32(len(text))
	var total float32 = 0.0

	for _, num := range histogram {
		floatNum := float32(num)
		total += (floatNum * (floatNum - 1))
	}
	return total / (length * (length - 1))
}
