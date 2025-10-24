package main

import (
	p "github.com/ulises39/portfolio-demo/pkg/portfolio"
)

func main() {
	portfolio1 := p.Portfolio{
		Name:          "Portfolio 1",
		Holdings:      make(map[string]p.Stock),
		AimingTargets: make(map[string]float64),
	}

	for i, stock := range p.DummyStockArr {
		portfolio1.AddStock(stock)
		portfolio1.SetTargetAllocation(stock.Ticker, p.DummyTargetArr[i])
	}
	portfolio1.Rebalance()

	randomStocks, randomTargets := p.GenerateRandomData()
	portfolio2 := p.NewPortfolio(randomStocks, randomTargets, "Portfolio 2")
	portfolio2.Rebalance()
}
