package models

// Query Params
type TradesParams struct{
	Book 	string 		`url:"book"`
	Marker	string		`url:"marker"`
	Sort	string		`url:"sort"` // 'asc' or 'desc'
	Limit	int			`url:"limit"` // max 100
}

type WithdrawParams struct{
	Currency	string	`json:"currency"`
	Amount		string	`json:"amount"`
	Address		string	`json:"address"`
	Tag			string	`json:"destination_tag"`
}

type TimeInForce string
const(
	GTC 	TimeInForce = "goodtillcancelled"
	FOR		TimeInForce = "fillorkill"
	IOC		TimeInForce = "immediateorcancel"
	PO		TimeInForce = "postonly"
)

type OrderSide string
const(
	Buy 	OrderSide = "buy"
	Sell 	OrderSide = "sell"
)

type PlaceOrderParams struct {
	Book 		string 		`json:"book"`
	Side		OrderSide	`json:"side"`
	TimeIF 		TimeInForce	`json:"time_in_force"`
	InternalID	string		`json:"client_id"`
}