package bitso

import (
	bitsoPrivate "github.com/grupokindynos/gobitso/private"
	bitsoPublic "github.com/grupokindynos/gobitso/public"
)

type Bitso struct {
	bitsoPrivate.BitsoPrivate
	bitsoPublic.BitsoPublic
}

func NewBitso(Url string) *Bitso{
	b := new(Bitso)
	b.Url = Url
	b.UrlPrivate = Url
	return b
}

// This enables private API functionality
func (b *Bitso) SetAuth(ApiKey string, ApiSecret string){
	b.ApiKey = ApiKey
	b.ApiSecret = ApiSecret
}