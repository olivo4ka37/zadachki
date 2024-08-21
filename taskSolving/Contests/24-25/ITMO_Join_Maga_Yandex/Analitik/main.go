package main

import (
	"fmt"
	"math/big"
)

const gridSize = 4

// Function to check if a 2x2 purple square exists in a given configuration
func has2x2PurpleSquare(grid [gridSize][gridSize]bool) bool {
	for i := 0; i < gridSize-1; i++ {
		for j := 0; j < gridSize-1; j++ {
			if grid[i][j] && grid[i][j+1] && grid[i+1][j] && grid[i+1][j+1] {
				return true
			}
		}
	}
	return false
}

func main() {
	totalConfigs := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(16), nil) // 2^16
	validConfigs := big.NewInt(0)

	// Iterate through all 2^16 possible configurations
	for i := 0; i < 1<<16; i++ {
		var grid [gridSize][gridSize]bool
		for j := 0; j < 16; j++ {
			grid[j/gridSize][j%gridSize] = (i & (1 << j)) != 0
		}
		if !has2x2PurpleSquare(grid) {
			validConfigs.Add(validConfigs, big.NewInt(1))
		}
	}

	m := validConfigs
	n := totalConfigs
	gcd := new(big.Int).GCD(nil, nil, m, n)

	m.Div(m, gcd)
	n.Div(n, gcd)

	mnSum := new(big.Int).Add(m, n)
	fmt.Println(m, n)
	fmt.Printf("Ответ: %s\n", mnSum.String())
}
