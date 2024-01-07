// in a package, function start with main function in package main (PSVM in java)
package main

// multiple import
import (
	"fmt"
	"math"
)

func printPi() {
	// in go, exported func / variable first letter capitialized (public)
	// math.pi is undefined
	fmt.Println(math.Pi)
}

// take 2 parameter both int
// also can written as func sum2(x int, y int)
func sum2(x, y int) int {
	return x + y
}

// take variant number of input(like a slice)
func sumN(nums ...int) int {
	total := 0
	// range return (index, number)
	for i, n := range nums {
		fmt.Printf("Index: %d, number: %d\n", i, n)
		total += n
	}
	return total
}

func main() {
	printPi()
	fmt.Println(sum2(1, 2))
	fmt.Println(sumN(1, 3, 5, 6))
}
