package operations

import (
	"profit-earnings/src/types"
)

const (
	Buy              = "buy"
	Sell             = "sell"
	TaxThreshold     = 20000.00
	TaxDuePercentage = 0.2
)

type Tax struct {
	TaxValue types.RoundedFloat64 `json:"tax"`
}

type Earnings struct {
	transactions      []types.Transaction
	averageStockPrice float64
	profit            float64
	stocks            int
}

// GetIncomingTaxes creates a new instance of Earnings
// and returns the tax value due to each transaction
// as a slice of Tax
func GetIncomingTaxes(transactions []types.Transaction) []Tax {
	e := NewEarnings(transactions)
	return e.CalculateIncomingTaxes()
}

func NewEarnings(transactions []types.Transaction) Earnings {
	return Earnings{transactions: transactions}
}

// CalculateIncomingTaxes calculates the tax due
// for each transaction based on its operation type
// and returns a slice of Tax
func (e *Earnings) CalculateIncomingTaxes() []Tax {
	taxes := make([]Tax, len(e.transactions))
	for idx, t := range e.transactions {
		switch t.Operation {
		case Buy:
			taxes[idx] = e.taxOverBuyTransaction(t)
		case Sell:
			taxes[idx] = e.taxOverSellTransaction(t)
		}
	}
	return taxes
}

// taxOverBuyTransaction returns a tax due value for a buy operation
func (e *Earnings) taxOverBuyTransaction(t types.Transaction) Tax {
	e.averageStockPrice = ((float64(e.stocks) * e.averageStockPrice) + (float64(t.Quantity) * t.UnitCost)) / float64(e.stocks+t.Quantity)
	e.stocks += t.Quantity
	return Tax{}
}

// taxOverSellTransaction returns a tax due value for a sell operation
func (e *Earnings) taxOverSellTransaction(t types.Transaction) Tax {
	var tax Tax
	totalSold := t.UnitCost * float64(t.Quantity)
	profit := (float64(e.stocks-t.Quantity)*e.averageStockPrice + totalSold) - float64(e.stocks)*e.averageStockPrice
	e.profit += profit
	if e.profit > 0 {
		if totalSold > TaxThreshold {
			tax.TaxValue = types.RoundedFloat64(e.profit * TaxDuePercentage)
		}
		e.profit = 0
	}
	e.stocks -= t.Quantity
	return tax
}
