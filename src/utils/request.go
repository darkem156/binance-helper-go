package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (client *Client) SendPublicRequest(request string, params map[string]string) string {
	response, _ := json.Marshal(params)
	return string(response)
}

func (client *Client) SendSignedRequest(uri string, params map[string]string, method string) string {
	queryParams := map[string]string{
		"recvWindow": "5000",
		"timestamp":  fmt.Sprintf("%d", time.Now().UnixNano()/1000000),
	}

	for key := range params {
		queryParams[key] = params[key]
	}

	urlValues := url.Values{}

	for key, value := range queryParams {
		urlValues.Add(key, value)
	}

	queryString := urlValues.Encode()

	hash := hmac.New(sha256.New, []byte(client.SecretKey))
	io.WriteString(hash, string(queryString))

	signature := hex.EncodeToString(hash.Sum(nil))

	println(signature)

	finalURL := uri + "?" + queryString + "&signature=" + signature

	var httpClient http.Client

	request, err := http.NewRequest(method, finalURL, nil)
	if err != nil {
		return err.Error()
	}

	request.Header.Add("X-MBX-APIKEY", client.ApiKey)
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	response, err := httpClient.Do(request)
	if err != nil {
		return err.Error()
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}
