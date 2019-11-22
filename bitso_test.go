package bitso

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/grupokindynos/gobitso/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const BitsoUrl = "https://api.bitso.com"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestAddressGeneration(t *testing.T) {
	b := NewBitso(BitsoUrl)
	b.SetAuth(os.Getenv("BITSO_API_KEY"), os.Getenv("BITSO_API_SECRET"))
	params := models.DestinationParams{
		FundCurrency: "btc",
	}
	address, err := b.FundingDestination(params)
	assert.Nil(t, err)
	assert.Equal(t, true, address.Success)
	assert.Equal(t, "Bitcoin address", address.Payload.AccountIdentifierName)
}

func TestPrivateApiAccess(t *testing.T) {
	b := NewBitso(BitsoUrl)
	_, err := b.Balances()
	assert.IsType(t, &models.NoCredentials{}, err)
}

func TestAvailableBooks(t *testing.T) {
	b := NewBitso(BitsoUrl)
	res, err := b.AvailableBooks()
	assert.Nil(t, err)
	assert.Equal(t, true, res.Success)
	// fmt.Println(res.Payload)
}

// Test Withdrawal
func TestWithdrawals(t *testing.T) {
	b := NewBitso(BitsoUrl)
	b.SetAuth(os.Getenv("BITSO_API_KEY"), os.Getenv("BITSO_API_SECRET"))
	params := models.WithdrawParams{
		Currency: "litecoin",
		Amount:   "0.000",
		Address:  "MQvNy1m7UZVfmeyAQEeYLYr9uJDwkAh898",
		Tag:      "Bitso API Unit Test",
	}
	_, err := b.Withdraw(params)
	assert.Nil(t, err)
	// fmt.Println("Test Response: ", res)
}

// Tests Private Api
func TestBalances(t *testing.T) {
	time.Sleep(3 * time.Second)
	b := NewBitso(BitsoUrl)
	b.SetAuth(os.Getenv("BITSO_API_KEY"), os.Getenv("BITSO_API_SECRET"))
	res, err := b.Balances()
	assert.Nil(t, err)
	assert.IsType(t, res, models.BalancesResponse{})
	assert.Equal(t, true, res.Success)
	fmt.Println(res)
}

// Tests Public API
func TestTrades(t *testing.T) {
	b := NewBitso(BitsoUrl)
	res, err := b.Trades(models.TradesParams{
		Book: "btc_mxn",
	})
	assert.Nil(t, err)
	assert.Equal(t, true, res.Success)
	assert.Equal(t, 25, len(res.Payload))
	assert.IsType(t, res, models.TradeResponse{})
}

func TestOrderPlacement(t *testing.T) {
	b := NewBitso(BitsoUrl)
	b.SetAuth(os.Getenv("BITSO_API_KEY"), os.Getenv("BITSO_API_SECRET"))
	_, err := b.PlaceOrder(models.PlaceOrderParams{
		Book:       "btc_mxn",
		Side:       "buy",
		TimeIF:     models.GTC,
		InternalID: "testing-order",
	})
	assert.Nil(t, err)
}

func TestTicker(t *testing.T) {
	b := NewBitso(BitsoUrl)
	_, err := b.Ticker(models.TickerParams{
		Book: "btc_mxn",
	})
	assert.Nil(t, err)
}
