package tools

func ScaleValue(value uint16, inMin float64, inMax float64, outMin float64, outMax float64) uint16 {
	valFloat := float64(value)
	result := (valFloat-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
	return uint16(result)
}
