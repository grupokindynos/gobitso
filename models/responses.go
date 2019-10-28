package models

import "time"

type TradeResponse struct {
	Success bool `json:"success"`
	Payload []OrderResponse `json:"payload"`
}

type OrderResponse struct {
	Book      string    `json:"book"`
	CreatedAt string	`json:"created_at"`
	Amount    string    `json:"amount"`
	MakerSide string    `json:"maker_side"`
	Price     string    `json:"price"`
	Tid       int       `json:"tid"`
}

type BalancesResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Balances []BalanceResponse `json:"balances"`
	} `json:"payload"`
}

type BalanceResponse struct {
	Currency  string `json:"currency"`
	Total     string `json:"total"`
	Locked    string `json:"locked"`
	Available string `json:"available"`
}

type WithdrawResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Wid       string    `json:"wid"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		Currency  string    `json:"currency"`
		Method    string    `json:"method"`
		Amount    string    `json:"amount"`
		Details   struct {
			WithdrawalAddress string      `json:"withdrawal_address"`
			TxHash            interface{} `json:"tx_hash"`
		} `json:"details"`
	} `json:"payload"`
}