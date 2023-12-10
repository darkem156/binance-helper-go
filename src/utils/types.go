package utils

type Client struct {
	ApiKey       string
	SecretKey    string
	BaseEndpoint string
}

type Klines struct {
	OpenTime                 []int64
	Open                     []float64
	High                     []float64
	Low                      []float64
	Close                    []float64
	Volume                   []float64
	CloseTime                []int64
	QuoteAssetVolume         []float64
	NumberOfTrades           []int64
	TakerBuyBaseAssetVolume  []float64
	TakerBuyQuoteAssetVolume []float64
}
