package util

import (
	"strings"
	"time"

	"github.com/mrod502/ddb"
	"github.com/shopspring/decimal"
)

type Quotable interface {
	ToCandle() Candle
	ddb.Indexer
}

//Candle - generic candle struct for database storage
type Candle struct {
	Begin  time.Time
	End    time.Time
	Open   decimal.Decimal
	High   decimal.Decimal
	Low    decimal.Decimal
	Close  decimal.Decimal
	Volume decimal.Decimal
	Symbol string
}

type Fill struct {
	FillID    string
	Exchange  rune
	Side      rune
	Volume    decimal.Decimal
	Timestamp time.Time
	AvgPx     decimal.Decimal
}

func (f Fill) OrderID() string {
	return strings.Split(f.FillID, string(IDSep))[0]
}
