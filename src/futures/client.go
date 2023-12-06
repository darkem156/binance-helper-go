package futures

import (
	"github.com/darkem156/binance-helper-go/src/utils"
)

func NewFuturesClient(apiKey, secretKey string) *utils.Client {
	return &utils.Client{ApiKey: apiKey, SecretKey: secretKey}
}
