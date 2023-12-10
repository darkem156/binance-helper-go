package futures

import (
	"github.com/darkem156/binance-helper-go/src/utils"
)

type FuturesClient struct {
	utils.Client
}

func NewFuturesClient(apiKey, secretKey string) *FuturesClient {
	return &FuturesClient{utils.Client{ApiKey: apiKey, SecretKey: secretKey, BaseEndpoint: "https://fapi.binance.com"}}
}
