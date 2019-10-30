package bitso_private

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/grupokindynos/gobitso/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type BitsoPrivate struct {
	ApiKey			string
	ApiSecret 		string
	UrlPrivate		string
}

func (b *BitsoPrivate) Balances() (models.BalancesResponse, error) {
	var balancesResp models.BalancesResponse
	data, err := b.PrivateRequest("/v3/balance", http.MethodGet, nil, nil)
	if err != nil {
		return balancesResp, err
	}
	err = json.Unmarshal(data, &balancesResp)
	if err != nil {
		return balancesResp, err
	}
	return balancesResp, nil
}

func(b *BitsoPrivate) Withdraw(params models.WithdrawParams) (models.WithdrawResponse, error) {
	var withdrawInfo models.WithdrawResponse
	byteParams, err := json.Marshal(params)
	if err != nil {
		return withdrawInfo, err
	}
	data, err := b.PrivateRequest("/v3/" + params.Currency + "_withdrawal", http.MethodPost, byteParams, nil)

	err = json.Unmarshal(data, &withdrawInfo)
	if err != nil {
		return models.WithdrawResponse{}, err
	}
	return withdrawInfo, nil
}

func (b *BitsoPrivate) PrivateRequest(url string, method string, params []byte, queryParams interface{}) ([]byte, error) {
	err := b.validateCredentials()
	if  err != nil {
		return []byte{}, err
	}

	var arr []byte
	client := &http.Client{}

	if method == http.MethodGet {
		key, nonce, signature := getSigningData(b.ApiKey, b.ApiSecret, method, url, params)
		var authH = fmt.Sprintf("Bitso %s:%s:%s", key, nonce, signature)
		req, err := http.NewRequest(method, b.UrlPrivate + url, nil)
		if err != nil {
			return arr, err
		}
		val, err := query.Values(queryParams)
		if err != nil {
			return arr, err
		}
		// fmt.Print("Encoded Values: ", val.Encode())
		q := req.URL.RawQuery
		req.URL.RawQuery = q + val.Encode()

		// Add Headers
		req.Header.Add("Authorization", authH)
		// Perform Request
		res, err := client.Do(req)

		if res != nil {
			defer res.Body.Close()
		}
		if err != nil {
			return arr, err
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return arr, err
		}
		return data, nil
	} else {
		req, err := http.NewRequest(method, b.UrlPrivate + url, bytes.NewBuffer(params))
		if err != nil {
			return arr, err
		}

		key, nonce, signature := getSigningData(b.ApiKey, b.ApiSecret, method, url, params)
		var authH = fmt.Sprintf("Bitso %s:%s:%s", key, nonce, signature)
		req.Header.Add("Authorization", authH)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if res != nil {
			defer res.Body.Close()
		}
		if err != nil {
			return arr, err
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return arr, err
		}
		return data, nil
	}
}

func getSigningData(key string, apiSecret string, method string, url string, params []byte) (string, string, string){
	var nonce = strconv.FormatInt(time.Now().Unix(), 10)
	var authHeather =  nonce + method + url + string(params)
	var signedPayload = hmac.New(sha256.New, []byte(apiSecret))
	signedPayload.Write([]byte(authHeather))
	var signature = hex.EncodeToString(signedPayload.Sum(nil))
	return key, nonce, signature
}

func (b *BitsoPrivate)validateCredentials() (error){
	if b.ApiKey == "" || b.ApiSecret == "" {
		return &models.NoCredentials{}
	}
	return nil
}