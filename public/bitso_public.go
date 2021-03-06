package bitso_public

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/grupokindynos/gobitso/models"
	"io/ioutil"
	"net/http"
)

type BitsoPublic struct {
	Url string
}

// Methods
func (b *BitsoPublic) Trades(params models.TradesParams) (models.TradeResponse, error) {
	var tradeResp models.TradeResponse
	data, err := b.PublicRequest("/v3/trades", http.MethodGet, nil, params)
	if err != nil {
		return tradeResp, err
	}
	err = json.Unmarshal(data, &tradeResp)
	if err != nil {
		return tradeResp, err
	}
	return tradeResp, nil
}

func (b *BitsoPublic) AvailableBooks() (models.BooksResponse, error) {
	var availableBooksResp models.BooksResponse
	data, err := b.PublicRequest("/v3/available_books", http.MethodGet, nil, nil)
	if err != nil {
		return availableBooksResp, err
	}
	err = json.Unmarshal(data, &availableBooksResp)
	return availableBooksResp, nil
}

func (b *BitsoPublic) Ticker(params models.TickerParams) (models.TickerResponse, error) {
	var tickerResp models.TickerResponse
	data, err := b.PublicRequest("/v3/ticker", http.MethodGet, nil, params)
	if err != nil {
		return tickerResp, err
	}
	err = json.Unmarshal(data, &tickerResp)
	return tickerResp, nil
}

func (b *BitsoPublic) PublicRequest(url string, method string, params []byte, queryParams interface{}) ([]byte, error) {
	var arr []byte
	client := &http.Client{}

	if method == http.MethodGet {
		req, err := http.NewRequest(method, b.Url+url, nil)
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
		// fmt.Println(q + req.URL.RawQuery)
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
		req, err := http.NewRequest(method, b.Url+url, bytes.NewBuffer(params))
		if err != nil {
			return arr, err
		}
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
