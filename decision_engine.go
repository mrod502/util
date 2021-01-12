package util

import (
	"fmt"
	"sync"
)

//ThinkerConfig - idk some config settngs
type ThinkerConfig struct {
	stuff []int
}

//Thinker - anything that can decide what to think about a symbol (must be its own isolated process)
type Thinker interface {
	Thoughts(string) Indicator
	Config() ThinkerConfig
	TakeQuotes(Quotable)
}

//DecisionEngine - da brain
type DecisionEngine struct {
	Workers      []*Worker
	router       map[string]int
	routerLock   *sync.RWMutex
	strats       *[]TradingStrategy
	decisionChan chan TradeSignal
	workerCtr    int
	numWorkers   int
	//	otherBrain Thinker
	//brainConfig ThinkerConfig
}

//NewDecisionEngine - return an initialized DecisionEngine
func NewDecisionEngine(numWorkers int, strats *[]TradingStrategy, decisionChan chan TradeSignal) *DecisionEngine {
	var d = &DecisionEngine{
		Workers:      make([]*Worker, numWorkers),
		router:       make(map[string]int),
		routerLock:   new(sync.RWMutex),
		decisionChan: decisionChan,
		numWorkers:   numWorkers,
	}

	for i := 0; i < numWorkers; i++ {
		d.Workers[i] = NewWorker(d.strats, 60)
		go d.Workers[i].StreamTrades(d.decisionChan)
	}
	return d
}

func (d *DecisionEngine) assignWorker(symbol string) {
	if len(d.Workers) == 0 {
		panic(fmt.Errorf("no workers %+v", d))
	}

	d.routerLock.Lock()
	defer d.routerLock.Unlock()

	_, ok := d.router[symbol]
	if ok {
		return
	}

	d.router[symbol] = d.workerCtr
	d.workerCtr = (d.workerCtr + 1) % d.numWorkers

}

func (d *DecisionEngine) findRoute(c Candle) int {
	d.routerLock.RLock()
	defer d.routerLock.RUnlock()
	if v, ok := d.router[c.Symbol]; ok {
		return v
	}
	return -1
}

//RouteQuotes  - route quotes to the workers, which in turn do things with the quotes
func (d *DecisionEngine) RouteQuotes(ch chan Quotable) {

	for {
		q := <-ch
		c := q.ToCandle()
		var ix int
		if ix = d.findRoute(c); ix == -1 {
			d.assignWorker(c.Symbol)
			ix = d.findRoute(c)
		}

		d.Workers[ix].push(c)

	}
}
