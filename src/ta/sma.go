package ta

func Sma(values []float64, period int) []float64 {
	var result []float64
	var sum float64
	for i := 0; i < len(values); i++ {
		sum += values[i]
		if i >= period {
			sum -= values[i-period]
			result = append(result, sum/float64(period))
		}
	}
	return result
}
