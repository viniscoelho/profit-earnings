package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"profit-earnings/src/operations"
	"profit-earnings/src/types"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		input := processInput(scanner)
		if len(input) == 0 {
			break
		}

		transactions := make([]types.Transaction, 0)
		err := json.Unmarshal([]byte(input), &transactions)
		if err != nil {
			log.Fatal(err)
		}

		taxes := operations.GetIncomingTaxes(transactions)
		taxesBytes, err := json.Marshal(taxes)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", taxesBytes)
	}
}

// processInput reads input from os.Stdin until it reaches a ']' rune
// or an empty line, suggesting end of current set of transactions
// and end of input respectively.
func processInput(scanner *bufio.Scanner) string {
	var input string
	for {
		scanner.Scan()
		line := scanner.Text()
		input += line
		if len(line) == 0 || line[len(line)-1] == ']' {
			break
		}
	}
	return input
}
