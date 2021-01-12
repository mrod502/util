package util

import (
	"github.com/shopspring/decimal"
)

//RSI - calculate RSI
func RSI(candles []Candle, p CandlePart) (rsi decimal.Decimal) {

	if len(candles) < 15 {
		rsi = decimal.New(5, 1)
		return
	}
	//get last 15
	candles = candles[len(candles)-15:]

	var avgGain = decimal.Zero
	var avgLoss = decimal.Zero

	ps := candlePartSlice(candles, p)

	ps = diffPercent(ps)
	pos, neg := splitPosNeg(ps)

	//handle short array (pos)
	if len(pos) > 1 {
		avgGain = decimal.Avg(pos[0], pos[1:]...)
	} else if len(pos) == 1 {
		avgGain = pos[0]
	}

	//handle short array (neg)
	if len(neg) > 1 {
		avgLoss = decimal.Avg(neg[0], neg[1:]...)
	} else if len(neg) == 1 {
		avgGain = neg[0]
	}

	// handle zeros
	if avgLoss.Equal(decimal.Zero) {
		if avgGain.GreaterThan(decimal.Zero) {
			rsi = hundred
			return
		}
		rsi = decimal.New(5, 1)
		return
	}

	rsi = hundred.Sub(hundred.Div(avgGain.Div(avgLoss).Add(one)))

	return

}
