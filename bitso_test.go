package bitso

import (
	"fmt"
	"github.com/grupokindynos/gobitso/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Test Withdrawal
func TestWithdrawals(t *testing.T) {
	b := NewBitso("https://api.bitso.com")
	b.SetAuth(os.Getenv("BITSO_API_KEY"), os.Getenv("BITSO_API_SECRET"))
	params := models.WithdrawParams{
		Currency: "litecoin",
		Amount:   "20",
		Address:  "MQvNy1m7UZVfmeyAQEeYLYr9uJDwkAh898",
	}
	res, err := b.Withdraw("litecoin", params)
	assert.Nil(t, err)
	fmt.Println("Test Response: ", res)
}
// Tests Private Api
func TestBalances(t *testing.T) {
	b := NewBitso("https://api.bitso.com")
	b.SetAuth(os.Getenv("BITSO_API_KEY"), os.Getenv("BITSO_API_SECRET"))
	res, err := b.Balances()
	assert.Nil(t, err)
	assert.IsType(t, res, models.BalancesResponse{})
	fmt.Println("TestBalances: ", res.Payload.Balances)
}

// Tests Public API
func TestTrades(t *testing.T) {
	b := NewBitso("https://api.bitso.com")
	res, err := b.Trades("btc_mxn")
	assert.Nil(t, err)
	assert.Equal(t, 25, len(res.Payload))
	assert.IsType(t, res, models.TradeResponse{})
}
