package utils

import (
	"fmt"
	"strconv"
)

func (client *Client) Klines(symbol string, interval string, limit int) Klines {
	lines, err := client.SendPublicRequest("/fapi/v1/klines", map[string]string{"symbol": symbol, "interval": interval, "limit": fmt.Sprint(limit)})

	if err != nil {
		fmt.Println(err)
		return Klines{}
	}

	klines := Klines{}

	for _, line := range lines.([]interface{}) {
		line := line.([]interface{})
		klines.OpenTime = append(klines.OpenTime, int64(line[0].(float64)))
		open, _ := strconv.ParseFloat(line[1].(string), 64)
		klines.Open = append([]float64{open}, klines.Open...)
		high, _ := strconv.ParseFloat(line[2].(string), 64)
		klines.High = append([]float64{high}, klines.High...)
		low, _ := strconv.ParseFloat(line[3].(string), 64)
		klines.Low = append([]float64{low}, klines.Low...)
		close, _ := strconv.ParseFloat(line[4].(string), 64)
		klines.Close = append([]float64{close}, klines.Close...)
		volume, _ := strconv.ParseFloat(line[5].(string), 64)
		klines.Volume = append([]float64{volume}, klines.Volume...)
		klines.CloseTime = append(klines.CloseTime, int64(line[6].(float64)))
		quoteAssetVolume, _ := strconv.ParseFloat(line[7].(string), 64)
		klines.QuoteAssetVolume = append([]float64{quoteAssetVolume}, klines.QuoteAssetVolume...)
		klines.NumberOfTrades = append(klines.NumberOfTrades, int64(line[8].(float64)))
		takerBuyBaseAssetVolume, _ := strconv.ParseFloat(line[9].(string), 64)
		klines.TakerBuyBaseAssetVolume = append([]float64{takerBuyBaseAssetVolume}, klines.TakerBuyBaseAssetVolume...)
		takerBuyQuoteAssetVolume, _ := strconv.ParseFloat(line[10].(string), 64)
		klines.TakerBuyQuoteAssetVolume = append([]float64{takerBuyQuoteAssetVolume}, klines.TakerBuyQuoteAssetVolume...)
	}

	return klines
}
