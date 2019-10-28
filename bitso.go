package bitso

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/gookit/color"
	"github.com/grupokindynos/gobitso/models"
	bitso_private "github.com/grupokindynos/gobitso/private"
	bitso_public "github.com/grupokindynos/gobitso/public"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Bitso struct {
	bitso_private.BitsoPrivate
	bitso_public.BitsoPublic
	Url 		string 	`json:"bitso_url"`
}

func NewBitso(Url string) *Bitso{
	b := new(Bitso)
	b.Url = Url
	return b
}

func (b *Bitso) SetAuth(ApiKey string, ApiSecret string){
	b.ApiKey = ApiKey
	b.ApiSecret = ApiSecret
}

func (b *Bitso)GetBalance() (bool, error) {

	data, err := b.PrivateRequest("/v3/balance", http.MethodGet, nil, nil)
	if err != nil {
		log.Println("err", err)
		// return tradeResp, err
	}

	log.Println(string(data))
	panic("not implemented yet")
}

func (b *Bitso) GetTrades(market string) (models.TradeResponse, error) {
	var tradeResp models.TradeResponse
	params := models.TradesParams{
		Book: market,
	}
	data, err := b.PublicRequest("/v3/trades", http.MethodGet, nil, params)
	if err != nil {
		return tradeResp, err
	}
	err = json.Unmarshal(data, &tradeResp)
	if err != nil {
		return tradeResp, err
	}
	fmt.Println("Response: ", tradeResp)
	return tradeResp, nil
}

func SignRequest() (string, error){
	return "", nil
}

func (b *Bitso)PublicRequest(url string, method string, params []byte, queryParams interface{}) ([]byte, error) {
	var arr []byte
	client := &http.Client{}

	if method == http.MethodGet {
		req, err := http.NewRequest(method, b.Url + url, nil)
		if err != nil {
			return arr, err
		}
		val, err := query.Values(queryParams)
		if err != nil {
			return arr, err
		}
		fmt.Print("Encoded Values: ", val.Encode())
		q := req.URL.RawQuery
		req.URL.RawQuery = q + val.Encode()
		fmt.Println(q + req.URL.RawQuery)
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
		req, err := http.NewRequest(method, b.Url + url, bytes.NewBuffer(params))
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

func (b *Bitso) PrivateRequest(url string, method string, params []byte, queryParams interface{}) ([]byte, error) {
	var arr []byte
	client := &http.Client{}

	var key = b.ApiKey
	var nonce = strconv.FormatInt(time.Now().Unix(), 10)
	var authHeather =  nonce + method + url + string(params)
	var signedPayload = hmac.New(sha256.New, []byte(b.ApiSecret))
	signedPayload.Write([]byte(authHeather))
	var signature = hex.EncodeToString(signedPayload.Sum(nil))
	var authH = fmt.Sprintf("Bitso %s:%s:%s", key, nonce, signature)


	log.Println(hex.EncodeToString(signedPayload.Sum(nil)))
	color.Info.Tips(key)
	log.Println(authH)

	if method == http.MethodGet {
		req, err := http.NewRequest(method, b.Url + url, nil)
		if err != nil {
			return arr, err
		}
		val, err := query.Values(queryParams)
		if err != nil {
			return arr, err
		}
		fmt.Print("Encoded Values: ", val.Encode())
		q := req.URL.RawQuery
		req.URL.RawQuery = q + val.Encode()
		fmt.Println(q + req.URL.RawQuery)

		// Add Headers
		req.Header.Add("Authorization", authH)
		color.Error.Tips(authHeather)

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
	panic("not implemented yet")
}