package models

type TradesParams struct{
	Book 	string 		`url:"book"`
	Marker	string		`url:"marker"`
	Sort	string		`url:"sort"` // 'asc' or 'desc'
	Limit	int			`url:"limit"` // max 100
}

type WithdrawParams struct{
	Currency	string	`url:"currency"`
	Amount		float64	`url:"amount"`
	Address		string	`url:"address"`
	Tag			string	`url:"destination_tag"`
}
