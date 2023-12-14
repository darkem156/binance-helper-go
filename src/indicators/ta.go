package indicators

import "github.com/darkem156/binance-helper-go/src/utils"

type Kline struct {
	utils.Klines
}

func Ta(symbol, interval string, limit int, onlyClosed bool) Kline {
	client := utils.Client{ApiKey: "", SecretKey: "", BaseEndpoint: "https://fapi.binance.com"}
	return Kline{client.Klines(symbol, interval, limit, onlyClosed)}
}
