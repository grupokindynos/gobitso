package bitso

import (
	"bytes"
	"fmt"
	"github.com/grupokindynos/gobitso/models"
	"io/ioutil"
	"net/http"
	"github.com/google/go-querystring/query"
)

type Bitso struct {
	ApiKey		string 	`json:"api_key"`
	ApiSecret 	string 	`json:"api_secret"`
	Url 		string 	`json:"bitso_url"`
}

func NewBitso(ApiKey string, ApiSecret string, Url string) *Bitso{
	b := new(Bitso)
	b.ApiKey = ApiKey
	b.ApiSecret = ApiSecret
	b.Url = Url
	return b
}

func GetBalance() (bool, error) {
	panic("not implemented yet")
}

func (b *Bitso) GetTrades(market string, limit int) (bool, error) {
	params := models.TradesParams{
		Book:market,
	}
	data, err := PublicRequest(b.Url + "/trades", http.MethodGet, nil, params)
	if err != nil {
		return false, err
	}
	fmt.Println(string(data))
	return true, nil
}

func SignRequest() (string, error){
	return "", nil
}

func PublicRequest(url string, method string, params []byte, queryParams interface{}) ([]byte, error) {
	var arr []byte
	client := &http.Client{}

	if method == http.MethodGet {
		req, err := http.NewRequest(method, url, nil)
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
		req, err := http.NewRequest(method, url, bytes.NewBuffer(params))
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