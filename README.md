# Gobitso API

This is a golang wrapper implementation for the Bitso REST API, developed by Kindynos

## Usage

In order to use Bitso's Public API just instantiate the service.

```go
import(
    "github.com/grupokindynos/gobitso"
)
b := NewBitso("https://api.bitso.com")
trades, _ := b.Trades("btc_mxn")
fmt.Println(trades)
```

To add private functionality use the SetAuth
```go
import(
    "github.com/grupokindynos/gobitso"
)
b := NewBitso("https://api.bitso.com")
b.SetAuth(os.Getenv("BITSO_API_KEY"), os.Getenv("BITSO_API_SECRET"))
balances, err := b.Balances()
fmt.Println(balances)
```


## Colaborators
* Luis Correa (Kindynos)