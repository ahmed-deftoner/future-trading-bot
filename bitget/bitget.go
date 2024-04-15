package bitget

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"
	"strconv"
	"time"
)

type BitgetBatchOrderRequest struct {
	Symbol        string         `json:"symbol"`
	MarginCoin    string         `json:"marginCoin"`
	OrderDataList []OrderRequest `json:"orderDataList"`
}

type OrderRequest struct {
	Size      string `json:"size"`
	Side      string `json:"side"`
	OrderType string `json:"orderType"`
}

type BatchOrderResponse struct {
	Code        string    `json:"code"`
	Msg         string    `json:"msg"`
	RequestTime int64     `json:"requestTime"`
	Data        BatchData `json:"data"`
}

type BatchData struct {
	OrderInfo []BatchOrderInfo `json:"orderInfo"`
	Failure   []interface{}    `json:"failure"`
	Result    bool             `json:"result"`
}

type BatchOrderInfo struct {
	OrderID   string `json:"orderId"`
	ClientOid string `json:"clientOid"`
}

func GetBitgetServerTimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

func GenerateBitgetSignature(apiSecret string, method string, uri string, timestamp string, requestBody string) string {
	message := ""
	if method == "GET" {
		message = fmt.Sprintf("%s%s%s", timestamp, method, uri)
	} else if method == "POST" {
		message = fmt.Sprintf("%s%s%s%s", timestamp, method, uri, requestBody)
	}

	hmac := hmac.New(sha256.New, []byte(apiSecret))
	hmac.Write([]byte(message))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	return signature
}

func PlaceBitgetBatchOrder(apiKey string, secretKey string, passphrase string, order *BitgetBatchOrderRequest) (BatchOrderResponse, error) {
	host := "https://api.bitget.com"
	path := "/api/mix/v1/order/batch-orders"
	url := host + path

	method := "POST"
	client := &http.Client{}

	jsonVal, err := json.Marshal(order)
	if err != nil {
		return BatchOrderResponse{}, err
	}

	serverTime := GetBitgetServerTimeStamp()
	signature := GenerateBitgetSignature(secretKey, "POST", path, serverTime, string(jsonVal))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonVal))
	req.Header.Add("ACCESS-KEY", apiKey)
	req.Header.Add("ACCESS-SIGN", signature)
	req.Header.Add("ACCESS-TIMESTAMP", serverTime)
	req.Header.Add("ACCESS-PASSPHRASE", passphrase)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("local", "zh-CN")

	if err != nil {
		return BatchOrderResponse{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return BatchOrderResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return BatchOrderResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return BatchOrderResponse{}, errors.New("bitget batch order failed")
	}

	var batchResponse BatchOrderResponse

	err = json.Unmarshal(body, &batchResponse)
	if err != nil {
		return BatchOrderResponse{}, err
	}

	return batchResponse, nil
}
