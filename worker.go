package util

import (
	"sync"
	"time"
)

//Worker - store and aggregate quotes
type Worker struct {
	quotes       []Candle
	strats       *[]TradingStrategy
	numElements  uint
	capacity     uint
	first        uint
	m            *sync.RWMutex
	maxTimestamp time.Time
	inputChan    chan Quotable
}

//Push - send a quotable to the thing
func (q *Worker) Push(qble Quotable) {
	q.inputChan <- qble
}

//Push - update the queue
func (q *Worker) push(c Candle) {

	q.m.Lock()
	defer q.m.Unlock()

	if ix := q.search(c); ix > -1 {
		q.quotes[ix] = c
		return
	}
	if q.numElements == q.capacity {
		q.first = (q.first + 1) % q.capacity
	}

}

//will use this for searching quotes array
func (q Worker) search(c Candle) int {

	return -1
}

//StreamTrades - process
func (q *Worker) StreamTrades(out chan TradeSignal) {

	var candle Candle
	for {
		quotable := <-q.inputChan
		candle = quotable.ToCandle()
		q.push(candle)
	}
}

//NewWorker - create a quotequeue
func NewWorker(strats *[]TradingStrategy, capacity uint) *Worker {
	return &Worker{quotes: make([]Candle, capacity), strats: strats, capacity: capacity, numElements: 0, inputChan: make(chan Quotable, 128)}
}
