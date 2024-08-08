package convert

import "strconv"

func MustToF32(str string) float32 {
	v, _ := strconv.ParseFloat(str, 32)
	return float32(v)
}

func MustToF64(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}
