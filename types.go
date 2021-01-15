package util

import (
	"strings"
	"time"

	"github.com/mrod502/ddb"
	"github.com/shopspring/decimal"
)

//Quotable - determines how to convert to a Candle
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

//Fill - an execution
type Fill struct {
	FillID    string
	Exchange  rune
	Side      rune
	Volume    decimal.Decimal
	Timestamp time.Time
	AvgPx     decimal.Decimal
}

//OrderID - the id of the order
func (f Fill) OrderID() string {
	return strings.Split(f.FillID, Sep)[0]
}

//Indicator - do we buy, sell, or hold?
type Indicator byte

//TradeSignal - tell someone what to trade
type TradeSignal struct {
	Side   Indicator
	Symbol string
}

//CandlePart - OHLC?
type CandlePart byte
