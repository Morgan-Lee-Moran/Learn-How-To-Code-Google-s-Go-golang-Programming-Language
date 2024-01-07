package main

import (
	"fmt"
	"math"
)

func positiveNegativeZero(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func positiveNegativeZero2(x int) int {
	// inline variable
	switch z := x; {
	case z > 0:
		return 1
	case z < 0:
		return -1
	default:
		return 0
	}
}

func sumPositive(x, y int) bool {
	// inline variable
	if z := x + y; z > 0 {
		return true
	}
	return false
}

func sumToX(start, end int) int {
	var total = 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total
}

// predefined loop condition variable
func sumToX2(start, end int) int {
	var total = 0
	for ; start <= end; start++ {
		total += start
	}
	return total
}

// while loop using for keyword
func sumToX3(start, end int) int {
	var total = 0
	for start <= end {
		total += start
		start++
	}
	return total
}

func Sqrt(x float64) float64 {
	z := 1.0
	// infinite loop
	for {
		z2 := z - (z*z-x)/(2*z)
		if math.Abs(z2-z) < 0.0000001 {
			return z2
		}
		z = z2
	}
}

func SqrtRecursion(x float64) float64 {
	return SqrtRecursionHelper(1.0, x)
}

func SqrtRecursionHelper(z, x float64) float64 {
	z2 := z - (z*z-x)/(2*z)
	if math.Abs(z2-z) < 0.0000001 {
		return z2
	} else {
		return SqrtRecursionHelper(z2, x)
	}
}

func main() {
	fmt.Println(positiveNegativeZero(0))
	fmt.Println(positiveNegativeZero2(-10))
	fmt.Println(sumToX(1, 10))
	fmt.Println(sumToX(1, 10) == sumToX2(1, 10))
	fmt.Println(sumToX(1, 10) == sumToX3(1, 10))
	fmt.Println(Sqrt(16))
	fmt.Println(SqrtRecursion(16))
}
