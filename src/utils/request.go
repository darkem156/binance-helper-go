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

func (client *Client) SendPublicRequest(endpoint string, params map[string]string) (interface{}, error) {
	queryParams := map[string]string{}
	for key := range params {
		queryParams[key] = params[key]
	}

	urlValues := url.Values{}

	for key, value := range queryParams {
		urlValues.Add(key, value)
	}

	queryString := urlValues.Encode()

	finalURL := client.BaseEndpoint + endpoint + "?" + queryString

	var httpClient http.Client

	request, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *Client) SendSignedRequest(endpoint string, params map[string]string, method string) (interface{}, error) {
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

	finalURL := client.BaseEndpoint + endpoint + "?" + queryString + "&signature=" + signature

	var httpClient http.Client

	request, err := http.NewRequest(method, finalURL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("X-MBX-APIKEY", client.ApiKey)
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
