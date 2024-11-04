package main

import (
	"fmt"
)

var memo map[[2]int]float64

func soupServings(n int) float64 {
	// Adjusting n for large values as probabilities converge to 1
	if n > 4800 {
		return 1.0
	}
	
	memo = make(map[[2]int]float64)
	return calculateProbability((n+24)/25, (n+24)/25)
}

func calculateProbability(a, b int) float64 {
	// Base cases
	if a <= 0 && b <= 0 {
		return 0.5 // Both are empty simultaneously
	}
	if a <= 0 {
		return 1.0 // A is empty first
	}
	if b <= 0 {
		return 0.0 // B is empty first
	}

	if val, exists := memo[[2]int{a, b}]; exists {
		return val
	}

	probability := 0.25 * (
		calculateProbability(a-4, b) +
		calculateProbability(a-3, b-1) +
		calculateProbability(a-2, b-2) +
		calculateProbability(a-1, b-3))
  
	memo[[2]int{a, b}] = probability

	return probability
}
