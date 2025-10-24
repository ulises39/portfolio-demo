package portfolio

import (
	"fmt"
)

func (p *Portfolio) TotalPortfolioValue() float64 {
	var total float64 = 0.0
	for _, stock := range p.Holdings {
		total += stock.Value()
	}
	return total
}

func (p *Portfolio) Rebalance() ([]Stock, []Stock) {
	outputStr := fmt.Sprintln("Rebalancing ", p.Name, "...")
	outputStr += p.ToString()

	var stocksToBuy []Stock
	var stocksToSell []Stock

	// 1. Calculate the total portfolio value(tpv)
	tpv := p.TotalPortfolioValue()
	outputStr += fmt.Sprintln("\nTotal portfolio value:", tpv)

	// 2. Determine target value for each stock.
	for ticker, target := range p.AimingTargets {
		outputStr += fmt.Sprintln("Processing", ticker)
		targetValue := tpv * target
		outputStr += fmt.Sprintln("\tTarget allocation:", target, "; Target value:", targetValue)

		// 3. Determine the difference. Target value vs current value.
		stock := p.Holdings[ticker]
		difference := targetValue - stock.Value()
		outputStr += fmt.Sprintln("\tDifference:", difference, "(target value - current value) -> ", targetValue, "-", stock.Price, "*", stock.Shares)

		// 4. Determine the number of shares to buy or sell.
		shares := int(difference / stock.Price)
		if shares > 0 { // Stock is underweight. Buy shares.
			outputStr += fmt.Sprintln("\t---> +  Buying", shares, "shares of", ticker)
			stocksToBuy = append(stocksToBuy, Stock{
				Ticker: ticker,
				Price:  stock.Price,
				Shares: shares,
			})
		} else if shares < 0 { // Stock is overweight. Sell shares.
			outputStr += fmt.Sprintln("\t---> - Selling", -1*shares, "shares of", ticker)
			stocksToSell = append(stocksToSell, Stock{
				Ticker: ticker,
				Price:  stock.Price,
				Shares: shares * -1,
			})
		} else { // Stock is at target allocation; as we need to sell or buy 0, there's no action needed. Do nothing.
			outputStr += fmt.Sprintln("\t---> 0 No action needed for", ticker)
		}
	}

	// 5. Return the stocks to buy and sell.
	outputStr += fmt.Sprintln("Rebalancing complete. Results:")
	outputStr += fmt.Sprintln("Stocks to buy:", stocksToBuy)
	outputStr += fmt.Sprintln("Stocks to sell:", stocksToSell)
	fmt.Println(outputStr)
	p.SaveReportFile(outputStr)
	return stocksToBuy, stocksToSell
}
