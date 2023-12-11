package indicators

import "github.com/darkem156/binance-helper-go/src/utils"

func Rma(values []float64, period int) []float64 {
	var ema []float64
	values = utils.Reverse(values)

	if len(values) < period {
		return ema
	}

	sum := 0.0
	for i := 0; i < period; i++ {
		sum += values[i]
	}
	ema = append(ema, sum/float64(period))

	multiplier := 1.0 / float64(period)
	for i := period; i < len(values); i++ {
		nextEMA := (values[i]-ema[len(ema)-1])*multiplier + ema[len(ema)-1]
		ema = append(ema, nextEMA)
	}

	return utils.Reverse(ema)
}

func (kline *Kline) Rma(source string, period int) []float64 {
	switch source {
	case "high":
		return Rma(kline.High, period)
	case "close":
		return Rma(kline.Close, period)
	case "low":
		return Rma(kline.Low, period)
	case "open":
		return Rma(kline.Open, period)
	default:
		return nil
	}
}
