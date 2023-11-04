package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

func SumIntOrFloats[K comparable, T int64 | float64](m map[K]T) T {
	var s T
	
	for _, v := range m {
		s += v
	}

	return s
}

func SumNumbers[K comparable, T Number](m map[K]T) T {
	var s T

	for _, v := range m {
		s += v
	}

	return s
} 

func main() {
	ints := map[string]int64 {
		"first": 34,
		"second": 12,
	}

	floats := map[string]float64 {
		"first": 35.98,
		"second": 26.99,
	}

	fmt.Printf("success Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Printf("success Generic Sums: %v and %v\n",
		SumIntOrFloats[string, int64](ints),
		SumIntOrFloats[string, float64](floats),
	)

	fmt.Printf("success Generic Sums, type parameters inferred: %v and %v\n",
		SumIntOrFloats(ints),
		SumIntOrFloats(floats),
	)

	fmt.Printf("success Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}