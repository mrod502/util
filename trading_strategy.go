package util

//TradingStrategy - trading strategy function template
type TradingStrategy func([]Candle, CandlePart) Indicator

//MajorityVote - apply strats to candles and poll to see the consensus
func MajorityVote(strats []TradingStrategy, candles []Candle, p CandlePart) Indicator {

	var b, s, h int32

	for _, strat := range strats {
		switch strat(candles, p) {
		case BUY:
			b++
		case SELL:
			s++
		case HOLD:
			h++
		default:
			h++
		}
	}

	if b > h {
		if b > s {
			return BUY
		}
		return SELL
	}
	if s > h {
		return SELL
	}
	return HOLD
}

func StreamSignals(quotes chan Candle, out chan TradeSignal, strats []TradingStrategy, candle CandlePart) {

	//var quotesMap = NewStore()

}
