package bitso_private

import (
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

func (b *BitsoPrivate) GetBalances() (models.BalancesResponse, error) {
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

func (b *BitsoPrivate) PrivateRequest(url string, method string, params []byte, queryParams interface{}) ([]byte, error) {
	var arr []byte
	client := &http.Client{}

	// Signing Data
	var key = b.ApiKey
	var nonce = strconv.FormatInt(time.Now().Unix(), 10)
	var authHeather =  nonce + method + url + string(params)
	var signedPayload = hmac.New(sha256.New, []byte(b.ApiSecret))
	signedPayload.Write([]byte(authHeather))
	var signature = hex.EncodeToString(signedPayload.Sum(nil))
	var authH = fmt.Sprintf("Bitso %s:%s:%s", key, nonce, signature)

	if method == http.MethodGet {
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
	}
	return []byte(nil), nil
}
