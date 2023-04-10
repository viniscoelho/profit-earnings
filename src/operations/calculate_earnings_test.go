package operations

import (
	"encoding/json"
	"profit-earnings/src/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	name     string
	input    string
	expected string
}

func TestScenarios(t *testing.T) {
	testCases := []TestCase{
		{
			name: "Test01",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 100},
			{"operation":"sell", "unit-cost":15.00, "quantity": 50},
			{"operation":"sell", "unit-cost":15.00, "quantity": 50}]`,
			expected: `[{"tax":0.00},{"tax":0.00},{"tax":0.00}]`,
		},
		{
			name: "Test02",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":20.00, "quantity": 5000},
			{"operation":"sell", "unit-cost":5.00, "quantity": 5000}]`,
			expected: `[{"tax":0.00},{"tax":10000.00},{"tax":0.00}]`,
		},
		{
			name: "Test03",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":5.00, "quantity": 5000},
			{"operation":"sell", "unit-cost":20.00, "quantity": 3000}]`,
			expected: `[{"tax":0.00},{"tax":0.00},{"tax":1000.00}]`,
		},
		{
			name: "Test04",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
			{"operation":"buy", "unit-cost":25.00, "quantity": 5000},
			{"operation":"sell", "unit-cost":15.00, "quantity": 10000}]`,
			expected: `[{"tax":0.00},{"tax":0.00},{"tax":0.00}]`,
		},
		{
			name: "Test05",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
			{"operation":"buy", "unit-cost":25.00, "quantity": 5000},
			{"operation":"sell", "unit-cost":15.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":25.00, "quantity": 5000}]`,
			expected: `[{"tax":0.00},{"tax":0.00},{"tax":0.00},{"tax":10000.00}]`,
		},
		{
			name: "Test06",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":2.00, "quantity": 5000},
			{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
			{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
			{"operation":"sell", "unit-cost":25.00, "quantity": 1000}]`,
			expected: `[{"tax":0.00},{"tax":0.00},{"tax":0.00},{"tax":0.00},{"tax":3000.00}]`,
		},
		{
			name: "Test07",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":2.00, "quantity": 5000},
			{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
			{"operation":"sell", "unit-cost":20.00, "quantity": 2000},
			{"operation":"sell", "unit-cost":25.00, "quantity": 1000},
			{"operation":"buy", "unit-cost":20.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":15.00, "quantity": 5000},
			{"operation":"sell", "unit-cost":30.00, "quantity": 4350},
			{"operation":"sell", "unit-cost":30.00, "quantity": 650}]`,
			expected: `[{"tax":0.00},{"tax":0.00},{"tax":0.00},{"tax":0.00},{"tax":3000.00},{"tax":0.00},{"tax":0.00},{"tax":3700.00},{"tax":0.00}]`,
		},
		{
			name: "Test08",
			input: `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":50.00, "quantity": 10000},
			{"operation":"buy", "unit-cost":20.00, "quantity": 10000},
			{"operation":"sell", "unit-cost":50.00, "quantity": 10000}]`,
			expected: `[{"tax":0.00},{"tax":80000.00},{"tax":0.00},{"tax":60000.00}]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transactions := make([]types.Transaction, 0)
			err := json.Unmarshal([]byte(tc.input), &transactions)
			assert.Nil(t, err)

			taxes := GetIncomingTaxes(transactions)
			taxesBytes, err := json.Marshal(taxes)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, string(taxesBytes))
		})
	}
}
