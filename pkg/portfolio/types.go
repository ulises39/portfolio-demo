package portfolio

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Stock struct {
	Ticker string
	Price  float64
	Shares int
}

// CurrentPrice updates the stock's price by recieving the last available price
func (s *Stock) CurrentPrice(price float64) {
	s.Price = price
}

func (s *Stock) GetPrice() float64 {
	return s.Price
}

// Value returns the total value of the stock
func (s *Stock) Value() float64 {
	return s.Price * float64(s.Shares)
}

type Portfolio struct {
	Name          string
	Holdings      map[string]Stock
	AimingTargets map[string]float64
}

func NewPortfolio(stocks []Stock, allocations []float64, name string) *Portfolio {

	p := &Portfolio{
		Name:          name,
		Holdings:      make(map[string]Stock),
		AimingTargets: make(map[string]float64),
	}

	for i, stock := range stocks {
		p.AddStock(stock)
		p.SetTargetAllocation(stock.Ticker, allocations[i])
	}

	return p
}

func (p *Portfolio) ToString() string {
	outputStr := "Holdings:\n"
	for _, stock := range p.Holdings {
		outputStr += fmt.Sprintln("\t[ Ticker:", stock.Ticker, "Shares:", stock.Shares, "Price:", stock.Price, "]")
	}
	outputStr += "Aiming Targets:\n"
	for ticker, allocation := range p.AimingTargets {
		outputStr += fmt.Sprintln("\t[ Ticker:", ticker, "Aim:", allocation, "]")
	}
	return outputStr
}

func (p *Portfolio) Print() {
	fmt.Println("Holdings:")
	for _, stock := range p.Holdings {
		fmt.Println("\t[ Ticker:", stock.Ticker, "Shares:", stock.Shares, "Price:", stock.Price, "]")
	}
	fmt.Println("Aiming Targets:")
	for ticker, allocation := range p.AimingTargets {
		fmt.Println("\t[ Ticker:", ticker, "Aim:", allocation, "]")
	}
}

func (p *Portfolio) AddStock(stock Stock) {
	p.Holdings[stock.Ticker] = stock
}

func (p *Portfolio) RemoveStock(ticker string) {
	delete(p.Holdings, ticker)
}

func (p *Portfolio) GetStock(ticker string) Stock {
	return p.Holdings[ticker]
}

func (p *Portfolio) GetTargetAllocation(ticker string) float64 {
	return p.AimingTargets[ticker]
}

func (p *Portfolio) SetTargetAllocation(ticker string, allocation float64) {
	p.AimingTargets[ticker] = allocation
}

func (p *Portfolio) SaveReportFile(outputStr string) error {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("rebalance_%s_%s.report.txt", p.Name, timestamp)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Warning: Could not create report file:", err)
		return err
	} else {
		defer file.Close()
	}

	var writer io.Writer
	if file != nil {
		writer = io.MultiWriter(os.Stdout, file)
	} else {
		writer = os.Stdout
	}

	_, err = fmt.Fprint(writer, outputStr)
	if err != nil {
		fmt.Println("Warning: Could not write to report file:", err)
		return err
	}
	return nil
}
