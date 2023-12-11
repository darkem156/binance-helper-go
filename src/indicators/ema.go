package indicators

import "github.com/darkem156/binance-helper-go/src/utils"

func Ema(values []float64, period int) []float64 {
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

	multiplier := 2.0 / float64(period+1)
	for i := period; i < len(values); i++ {
		nextEMA := (values[i]-ema[len(ema)-1])*multiplier + ema[len(ema)-1]
		ema = append(ema, nextEMA)
	}

	return utils.Reverse(ema)
}

func (kline *Kline) Ema(source string, period int) []float64 {
	switch source {
	case "high":
		return Ema(kline.High, period)
	case "close":
		return Ema(kline.Close, period)
	case "low":
		return Ema(kline.Low, period)
	case "open":
		return Ema(kline.Open, period)
	default:
		return nil
	}
}
