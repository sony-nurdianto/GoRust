package main

import (
	"fmt"
	"time"
)

// Lookup table for the original PopCount
var pc [256]byte

// Original PopCount using table lookup
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// New PopCount using bitwise shifting
func PopCountShift(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ { // Iterate through 64 bits
		count += int(x & 1) // Check the rightmost bit
		x >>= 1             // Shift bits to the right
	}
	return count
}

func init() {
	// Initialize the lookup table for PopCountTable
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	const testValue = 123456789 // Test number
	const iterations = 1000000  // Number of iterations for benchmarking

	// Benchmark the table-lookup version
	startTable := time.Now()
	for i := 0; i < iterations; i++ {
		PopCountTable(testValue)
	}
	elapsedTable := time.Since(startTable)

	// Benchmark the bitwise-shifting version
	startShift := time.Now()
	for i := 0; i < iterations; i++ {
		PopCountShift(testValue)
	}
	elapsedShift := time.Since(startShift)

	// Print results
	fmt.Printf("PopCountTable(%d) = %d\n", testValue, PopCountTable(testValue))
	fmt.Printf("PopCountShift(%d) = %d\n", testValue, PopCountShift(testValue))
	fmt.Printf("Table-lookup time: %v\n", elapsedTable)
	fmt.Printf("Bitwise-shifting time: %v\n", elapsedShift)
}
