# Portfolio Rebalancer Demo

This is a demo of a portfolio rebalancer. It can be used to rebalance a portfolio of stocks to a target allocation.

For now, is intended to be used only as a console application and with the internal examples.

## Requirements

- Go 1.22.4 or higher

## Features

- Rebalance a portfolio of stocks to a target allocation.
- Generate a report file with the results.
- Generate random data for testing. This includes:
  - Random stocks. Each stock has a random ticker (4 characters), random shares (between 1 and 20), and random price (between 1.00 and 100.00).
  - Random target allocations (based on the number of stocks). Ensuring the sum of the allocations is 1.0.

## Installation and Usage

```bash
git clone https://github.com/ulises39/portfolio-demo.git
cd portfolio-demo

# Run the demo
go run main.go
```

This should output in the console the calculation process, with the last lines being the results of wich stocks to buy and sell. Also, it will generate a report file in the current directory.

## Rebalance Process and Algorithm

The rebalance process is as follows:

1.  Calculate the total portfolio value (TPV).
    Sum the current market value of all holdings.

        TPV = Σ (Stock Shares × Stock Price)

2.  For each stock, calculate the target value (TV) based on the target allocation.

    Target Value_i = TPV × Target Allocation_i

3.  Calculate the difference between the target value and the current value.

    Difference_i = Target Value_i - (Stock Shares × Stock Price)

    - A **positive Difference** means the stock is **Underweight** (should **BUY** more).
    - A **negative Difference** means the stock is **Overweight** (should **SELL** some).

4.  Calculate the number of shares to buy or sell based on the difference.

    Shares to Buy/Sell_i = Difference_i / Stock Price_i

    > **Note:** Since this is a simple model, we can only buy/sell whole shares (i.e, integer values)

5.  Return the stocks to buy and sell.

