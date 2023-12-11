package utils

func Add(source []float64, value float64) []float64 {
	return append([]float64{value}, source...)
}

func Reverse(source []float64) []float64 {
	var result []float64
	for i := len(source) - 1; i >= 0; i-- {
		result = append(result, source[i])
	}
	return result
}
