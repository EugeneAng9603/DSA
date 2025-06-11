// 🧠 Strategy
// Compute net balances.

// Separate into debtors and creditors.

// Use a greedy algorithm to cancel debts with minimal transactions.

package main

import (
	"fmt"
)

type Transaction struct {
	from   int
	to     int
	amount int
}

func minimizeTransactions(expenses []int) []Transaction {
	n := len(expenses)
	total := 0
	for _, e := range expenses {
		total += e
	}
	avg := total / n

	// Step 1: Calculate net balance per user
	balance := make([]int, n)
	for i := 0; i < n; i++ {
		balance[i] = expenses[i] - avg
	}

	// Step 2: Greedy settle balances
	var transactions []Transaction
	i, j := 0, 0
	for {
		// Find first debtor (balance < 0) and creditor (balance > 0)
		for i < n && balance[i] >= 0 {
			i++
		}
		for j < n && balance[j] <= 0 {
			j++
		}
		if i >= n || j >= n {
			break
		}

		// Settle between i and j
		amount := min(-balance[i], balance[j])
		balance[i] += amount
		balance[j] -= amount

		transactions = append(transactions, Transaction{
			from:   i,
			to:     j,
			amount: amount,
		})
	}

	return transactions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 🔍 Usage Example
func main() {
	expenses := []int{15, 100, 50, 61, 37, 82, 33}
	txns := minimizeTransactions(expenses)

	for _, txn := range txns {
		fmt.Printf("User%d pays User%d: $%d\n", txn.from, txn.to, txn.amount)
	}
}
