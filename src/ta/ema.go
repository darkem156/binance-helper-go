package ta

func Ema(values []float64, period int) []float64 {
	var ema []float64

	if len(values) < period {
		return ema
	}

	sum := 0.0
	for i := 0; i < period; i++ {
		sum += values[i]
	}
	ema = append(ema, sum/float64(period))

	multiplier := 2.0 / float64(period+1)
	for i := period; i < len(values); i++ {
		nextEMA := (values[i]-ema[len(ema)-1])*multiplier + ema[len(ema)-1]
		ema = append(ema, nextEMA)
	}

	return ema
}
