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
	//Tag			string	`json:"destination_tag"`
}
