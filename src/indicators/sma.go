package indicators

import "github.com/darkem156/binance-helper-go/src/utils"

func Sma(values []float64, period int) []float64 {
	values = utils.Reverse(values)
	var result []float64
	var sum float64
	for i := 0; i < len(values); i++ {
		sum += values[i]
		if i >= period {
			sum -= values[i-period]
			result = utils.Add(result, sum/float64(period))
		}
	}
	return result
}

func (kline *Kline) Sma(source string, period int) []float64 {
	switch source {
	case "high":
		return Sma(kline.High, period)
	case "close":
		return Sma(kline.Close, period)
	case "low":
		return Sma(kline.Low, period)
	case "open":
		return Sma(kline.Open, period)
	default:
		return nil
	}
}
