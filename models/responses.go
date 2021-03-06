package models

import "time"

type TradeResponse struct {
	Success bool            `json:"success"`
	Payload []OrderResponse `json:"payload"`
}

type TickerResponse struct {
	Success bool             `json:"success"`
	Payload BookInfoResponse `json:"payload"`
}

type BookInfoResponse struct {
	Book      string    `json:"book"`
	Volume    string    `json:"volume"`
	High      string    `json:"high"`
	Last      string    `json:"last"`
	Low       string    `json:"low"`
	Vwap      string    `json:"vwap"`
	Ask       string    `json:"ask"`
	Bid       string    `json:"bid"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderResponse struct {
	Book      string `json:"book"`
	CreatedAt string `json:"created_at"`
	Amount    string `json:"amount"`
	MakerSide string `json:"maker_side"`
	Price     string `json:"price"`
	Tid       int    `json:"tid"`
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

type BooksResponse struct {
	Success bool           `json:"success"`
	Payload []BookResponse `json:"payload"`
}

type BookResponse struct {
	Book          string `json:"book"`
	MinimumAmount string `json:"minimum_amount"`
	MaximumAmount string `json:"maximum_amount"`
	MinimumPrice  string `json:"minimum_price"`
	MaximumPrice  string `json:"maximum_price"`
	MinimumValue  string `json:"minimum_value"`
	MaximumValue  string `json:"maximum_value"`
}

type AccountInfoResponse struct {
	Success bool    `json:"success"`
	Payload Account `json:"payload"`
}

type Account struct {
	ClientID              string `json:"client_id"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Status                string `json:"status"`
	DailyLimit            string `json:"daily_limit"`
	MonthlyLimit          string `json:"monthly_limit"`
	DailyRemaining        string `json:"daily_remaining"`
	MonthlyRemaining      string `json:"monthly_remaining"`
	CashDepositAllowance  string `json:"cash_deposit_allowance"`
	CellphoneNumber       string `json:"cellphone_number"`
	CellphoneNumberStored string `json:"cellphone_number_stored"`
	EmailStored           string `json:"email_stored"`
	OfficialID            string `json:"official_id"`
	ProofOfResidency      string `json:"proof_of_residency"`
	SignedContract        string `json:"signed_contract"`
	OriginOfFunds         string `json:"origin_of_funds"`
}

type PlacedOrderResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Oid string `json:"oid"`
	} `json:"payload"`
}

type DestinationResponse struct {
	Success bool `json:"success"`
	Payload struct {
		AccountIdentifierName string `json:"account_identifier_name"`
		AccountIdentifier     string `json:"account_identifier"`
	} `json:"payload"`
}

type LookUpOrdersResponse struct {
	Success bool                  `json:"success"`
	Payload []LookUpOrderResponse `json:"payload"`
}
type LookUpOrderResponse struct {
	Book           string  `json:"book"`
	OriginalAmount string  `json:"original_amount"`
	UnfilledAmount float64 `json:"unfilled_amount"`
	OriginalValue  string  `json:"original_value"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Price          string  `json:"price"`
	Oid            string  `json:"oid"`
	Side           string  `json:"side"`
	Status         string  `json:"status"`
	Type           string  `json:"type"`
}

type UserTradesResponse struct {
	Success bool                `json:"success"`
	Payload []UserTradeResponse `json:"payload"`
}

type UserTradeResponse struct {
	Book         string `json:"book"`
	Major        string `json:"major"`
	CreatedAt    string `json:"created_at"`
	Minor        string `json:"minor"`
	FeesAmount   string `json:"fees_amount"`
	FeesCurrency string `json:"fees_currency"`
	Price        string `json:"price"`
	Tid          string `json:"tid"`
	Oid          string `json:"oid"`
	Side         string `json:"side"`
}

type BookFees struct {
	Book            string `json:"book"`
	TakerFeeDecimal string `json:"taker_fee_decimal"`
	TakerFeePercent string `json:"taker_fee_percent"`
	MakerFeeDecimal string `json:"maker_fee_decimal"`
	MakerFeePercent string `json:"maker_fee_percent"`
}

type UserFeeResponse struct {
	Fees           []BookFees        `json:"fees"`
	WithdrawalFees map[string]string `json:"withdrawal_fees"`
}

type UserFeesResponse struct {
	Success bool            `json:"success"`
	Payload UserFeeResponse `json:"payload"`
}

type FundingResponse struct {
	Success bool `json:"success"`
	Payload []struct {
		Fid       string    `json:"fid"`
		Status    string    `json:"status"`
		CreatedAt string 	`json:"created_at"`
		Currency  string    `json:"currency"`
		Method    string    `json:"method"`
		Amount    string    `json:"amount"`
		Details   FundingDetails `json:"details"`
	} `json:"payload"`
}

type FundingDetails struct {
	SenderName       string `json:"sender_name"`
	SenderBank       string `json:"sender_bank"`
	SenderClabe      string `json:"sender_clabe"`
	ReceiveClabe     string `json:"receive_clabe"`
	NumericReference string `json:"numeric_reference"`
	Concepto         string `json:"concepto"`
	ClaveRastreo     string `json:"clave_rastreo"`
	BeneficiaryName  string `json:"beneficiary_name"`
	FundingAddress string `json:"funding_address"`
	TxHash         string `json:"tx_hash"`
}
