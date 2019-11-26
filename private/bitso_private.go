package bitso_private

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/grupokindynos/gobitso/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type BitsoPrivate struct {
	ApiKey     string
	ApiSecret  string
	UrlPrivate string
}

func (b *BitsoPrivate) AccountStatus() (accountInfo models.AccountInfoResponse, err error) {
	/*
		Retrieves Bitso account's status.
	 */
	data, err := b.PrivateRequest("/v3/account_status", http.MethodGet, nil, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &accountInfo)
	if err != nil {
		return
	}
	return accountInfo, err
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

func (b *BitsoPrivate) FundingDestination(params models.DestinationParams) (models.DestinationResponse, error) {
	var destinationResp models.DestinationResponse
	data, err := b.PrivateRequest("/v3/funding_destination", http.MethodGet, nil, params)
	if err != nil {
		return destinationResp, err
	}
	err = json.Unmarshal(data, &destinationResp)
	if err != nil {
		return destinationResp, err
	}
	return destinationResp, nil
}

func (b *BitsoPrivate) PlaceOrder(params models.PlaceOrderParams) (models.PlacedOrderResponse, error) {
	var placedOrderResp models.PlacedOrderResponse
	byteParams, err := json.Marshal(params)
	if err != nil {
		return placedOrderResp, err
	}
	data, err := b.PrivateRequest("/v3/orders", http.MethodPost, byteParams, nil)
	if err != nil {
		return placedOrderResp, err
	}
	err = json.Unmarshal(data, &placedOrderResp)
	if err != nil {
		return placedOrderResp, err
	}
	return placedOrderResp, nil
}

func (b *BitsoPrivate) Withdraw(params models.WithdrawParams) (models.WithdrawResponse, error) {
	var withdrawInfo models.WithdrawResponse
	byteParams, err := json.Marshal(params)
	if err != nil {
		return withdrawInfo, err
	}
	data, err := b.PrivateRequest("/v3/"+params.Currency+"_withdrawal", http.MethodPost, byteParams, nil)

	err = json.Unmarshal(data, &withdrawInfo)
	if err != nil {
		return models.WithdrawResponse{}, err
	}
	return withdrawInfo, nil
}

func (b *BitsoPrivate) UserTrades(params models.UserTradesParams) (userTrades models.UserTradesResponse, err error) {
	data, err := b.PrivateRequest("/v3/user_trades", http.MethodGet, nil, params)
	fmt.Println("UserTradesData: ", string(data))
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &userTrades)
	if err != nil {
		return
	}
	return userTrades, err
}

func (b *BitsoPrivate) OrderTrades(params models.OrderTradesParams) (userTrades models.UserTradesResponse, err error) {
	data, err := b.PrivateRequest("/v3/order_trades/order_trades/" + params.Oid + "/", http.MethodGet, nil, nil)
	fmt.Println("OrderTrades: ", string(data))
	var errResponse models.ErrorResponse
	if err != nil {
		fmt.Println(err)
		return userTrades, err
	}
	err = json.Unmarshal(data, &userTrades)
	if err != nil {
		return userTrades, err
	}
	if userTrades.Success == false {
		err = json.Unmarshal(data, errResponse)
		if err != nil {
			return userTrades, errors.New(errResponse.Error.Message)
		}
	}
	return userTrades, err
}

func (b *BitsoPrivate) OpenOrders() {

}

func (b *BitsoPrivate) LookUpOrders(orders []string) (models.LookUpOrdersResponse, error) {
	/*
		Looks for an order status in a user's history

		// TODO Support for client_id
	 */
	fmt.Println()
	var orderStr string
	if len(orders) == 1 {
		orderStr = orders[0]
	} else {
		for i, order := range orders {
			if i == 0 {
				orderStr = order
			} else {
				orderStr += "-" + order
			}

		}
	}

	fmt.Println("LookUpOrdersUrlOrders: ", orderStr)
	var orderInfo models.LookUpOrdersResponse
	data, err := b.PrivateRequest("/v3/orders/" + orders[0] + "/", http.MethodGet, nil, nil)
	fmt.Println(string(data))
	err = json.Unmarshal(data, &orderInfo)
	if err != nil {
		return models.LookUpOrdersResponse{}, err
	}
	fmt.Println("Done Request")
	return orderInfo, nil
}

func (b *BitsoPrivate) PrivateRequest(url string, method string, params []byte, queryParams interface{}) ([]byte, error) {
	err := b.validateCredentials()
	if err != nil {
		return []byte{}, err
	}

	var arr []byte
	client := &http.Client{}

	if method == http.MethodGet {
		val, err := query.Values(queryParams)
		if err != nil {
			return arr, err
		}
		qParams := val.Encode()
		bRequestUrl := url
		if queryParams != nil {
			bRequestUrl = url + "?" + qParams
		}
		key, nonce, signature := getSigningData(b.ApiKey, b.ApiSecret, method, bRequestUrl, nil)
		var authH = fmt.Sprintf("Bitso %s:%s:%s", key, nonce, signature)
		req, err := http.NewRequest(method, b.UrlPrivate+url, nil)
		if err != nil {
			return arr, err
		}
		// fmt.Print("Encoded Values: ", val.Encode())
		q := req.URL.RawQuery
		req.URL.RawQuery = q + qParams

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
		req, err := http.NewRequest(method, b.UrlPrivate+url, bytes.NewBuffer(params))
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

func getSigningData(key string, apiSecret string, method string, url string, params []byte) (string, string, string) {
	var nonce = strconv.FormatInt(time.Now().Unix(), 10)
	var authHeather = nonce + method + url + string(params)
	var signedPayload = hmac.New(sha256.New, []byte(apiSecret))
	signedPayload.Write([]byte(authHeather))
	var signature = hex.EncodeToString(signedPayload.Sum(nil))
	fmt.Println("Nonce: ", nonce, url)
	return key, nonce, signature
}

func (b *BitsoPrivate) validateCredentials() error {
	if b.ApiKey == "" || b.ApiSecret == "" {
		return &models.NoCredentials{}
	}
	return nil
}
