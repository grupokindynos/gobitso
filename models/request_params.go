package models

type TradesParams struct{
	Book 	string 		`url:"book"`
	Marker	string		`url:"marker"`
	Sort	string		`url:"sort"` // 'asc' or 'desc'
	Limit	int			`url:"limit"` // max 100
}
