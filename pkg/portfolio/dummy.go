package portfolio

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var DummyStockArr []Stock = []Stock{
	{
		Ticker: "AAPL",
		Shares: 10,
		Price:  100.00,
	},
	{
		Ticker: "MSFT",
		Shares: 10,
		Price:  100.00,
	},
	{
		Ticker: "GOOG",
		Shares: 10,
		Price:  100.00,
	},
	{
		Ticker: "META",
		Shares: 10,
		Price:  100.00,
	},
	{
		Ticker: "AMZN",
		Shares: 10,
		Price:  100.00,
	},
}

// DummyTargetArr is an array of target allocations for the dummy stocks. It's
// length must match the length of DummyStockArr and the values must sum to 1.0
var DummyTargetArr []float64 = []float64{0.1, 0.3, 0.3, 0.1, 0.2}

func randomTicker(n int, randomGenerator *rand.Rand) string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[randomGenerator.Intn(len(letters))]
	}
	return string(b)
}

func generateRandomStockArr() []Stock {

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)

	// 1. Random count between 5 and 10
	n := randomGenerator.Intn(6) + 5 // [0, 5] + 5 => [5, 10]

	stocks := make([]Stock, 0, n)
	usedTickers := make(map[string]struct{})

	for len(stocks) < n {
		// 2. Unique 4-char ticker
		ticker := randomTicker(4, randomGenerator)
		if _, exists := usedTickers[ticker]; exists {
			continue
		}
		usedTickers[ticker] = struct{}{}

		// 3. Shares 1 - 20
		shares := randomGenerator.Intn(20) + 1

		// 4. Price 1.00 - 100.00 with two decimals
		price := math.Round((randomGenerator.Float64()*99.0+1.0)*100) / 100

		stocks = append(stocks, Stock{
			Ticker: ticker,
			Shares: shares,
			Price:  price,
		})
	}

	return stocks
}

func generateRandomTargetArr(stocks []Stock) []float64 {

	N := len(stocks)

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)

	upperBound := 1.0
	lowerBound := 0.01

	// Minimum percentage for any stock (0.01)
	const minAllocation = 0.01

	// 1. Generate random target allocations
	targetArr := make([]float64, N)
	for i := range N - 1 {
		fmt.Println("\nProcessing stock", i)
		fmt.Println("\tUpper bound:", upperBound)
		fmt.Println("\tLower bound:", lowerBound)
		// Calculate the minimum value that MUST be reserved for the remaining stocks (including the last one).
		// The number of stocks that still need an allocation is N - (i + 1)
		remainingToAllocate := N - (i + 1)
		fmt.Println("\tRemaining items to allocate:", remainingToAllocate)
		lowerBound = float64(remainingToAllocate) * minAllocation
		fmt.Println("\tRequired minimum:", lowerBound)

		// The maximum value we can allocate to the current stock is the remainder
		// minus the required minimum for the remaining stocks.
		maxAllocation := upperBound - lowerBound
		fmt.Println("\tMax available allocation:", maxAllocation)

		// We ensure the max allocation is at least the minAllocation
		if maxAllocation < minAllocation {
			maxAllocation = minAllocation
			fmt.Println("\tMax allocation is less than min allocation. Setting max allocation to min allocation.")
		}

		// Generate a random float between minAllocation and maxAllocation
		// Formula: rand.Float64() * (length of range) + minimum value
		target := randomGenerator.Float64()*(maxAllocation-minAllocation) + minAllocation

		// Force rounding to two decimal places
		target = math.Round(target*100) / 100
		fmt.Println("\tTarget before rounding:", target)

		// If rounding caused the target to exceed the maximum allowed for safety, cap it.
		if target > maxAllocation {
			target = maxAllocation
			fmt.Println("\tTarget exceeded max allocation. Setting target to max allocation.")
		}

		targetArr[i] = target
		upperBound -= target
		fmt.Println("\tUpper bound after allocation:", upperBound)

		// IMPORTANT: Re-round the remainingTarget after subtraction to avoid cumulative float error
		// when calculating the next maxAllocation and to ensure the remainder is two-decimal friendly.
		upperBound = math.Round(upperBound*100) / 100
		fmt.Println("\tUpper bound after rounding:", upperBound)
	}

	// The last stock takes the remaining amount, absorbing the cumulative rounding error.
	targetArr[N-1] = upperBound
	fmt.Println("\tLast stock allocation:", targetArr[N-1])

	// Verify sum (for debugging and demonstration)
	sum := 0.0
	for _, target := range targetArr {
		sum += target
	}
	fmt.Printf("\n--- Random Allocation Check ---\n")
	fmt.Printf("Generated Allocations (Sum: %.2f):\n", sum)
	for i, alloc := range targetArr {
		fmt.Printf("  Stock %d: %.2f\n", i+1, alloc)
	}

	// Because of the rounding and remainder absorption, this sum should be exactly 1.0
	if math.Abs(sum-1.0) > 0.000001 {
		fmt.Println("CRITICAL ERROR: Sum is not 1.0")
	}

	return targetArr
}

func GenerateRandomData() ([]Stock, []float64) {
	randomStocks := generateRandomStockArr()
	randomTargetArr := generateRandomTargetArr(randomStocks)
	return randomStocks, randomTargetArr
}
