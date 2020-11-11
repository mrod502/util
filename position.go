package util

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

//Position - track positions
type Position struct {
	PID            []byte
	Symbol         string
	Qty            decimal.Decimal
	Open           decimal.Decimal
	BuyPx          decimal.Decimal
	BuyLmt         decimal.Decimal
	Exchange       string
	SellPx         decimal.Decimal
	SellLmt        decimal.Decimal
	BCommission    decimal.Decimal
	SCommission    decimal.Decimal
	OpenTimestamp  time.Time
	CloseTimestamp time.Time
	Short          bool
	OrderIDs       [][]byte
}

//ID indexer
func (p Position) ID() []byte {
	return p.PID
}

//Type Indexer
func (p Position) Type() []byte {
	return []byte(TblPosition)
}

//Base36 encodes an int64 as a base36 string representation (for db keys)
func Base36(n int64) string {
	return strconv.FormatInt(n, 36)
}
