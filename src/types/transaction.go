package types

// Transaction handles the input of a single transaction operation
type Transaction struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}
