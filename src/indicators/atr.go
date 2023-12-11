package indicators

import (
	"math"

	"github.com/darkem156/binance-helper-go/src/utils"
)

func (kline *Kline) Atr(period int) []float64 {
	high := kline.High
	low := kline.Low
	close := kline.Close

	true_ranges := []float64{}

	for i := len(high) - 1; i >= 0; i-- {
		if i == len(high)-1 {
			true_ranges = utils.Add(true_ranges, high[i]-low[i])
		} else {
			true_ranges = utils.Add(true_ranges, math.Max(high[i]-low[i], math.Max(math.Abs(high[i]-close[i+1]), math.Abs(low[i]-close[i+1]))))
		}
	}

	atr := Rma(true_ranges, period)

	return atr
}
