package models

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
