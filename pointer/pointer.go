package main

import (
	"fmt"
	"strings"
)

func pointer() {
	var i int = 1
	var p *int = &i
	// read value
	fmt.Println("i = ", *p)
	// write value
	*p = 21
	fmt.Println("i = ", *p)
}

type Vertex struct {
	X int
	Y int
}

func accessStruct() {
	// assign 1 to x, y have default value 0
	v := Vertex{X: 1}
	p := &v
	p.X = 5
	fmt.Println(v)
}

func sliceOper(width, height int) {
	var s [][]int = make([][]int, height)
	for i := 0; i < 5; i++ {
		s[i] = make([]int, width)
	}
	fmt.Println(len(s))
	// assign value
	s[0][0] = 1
	fmt.Println(s[0][0])
	// default value 0
	fmt.Println(s[1][0])
	// append a new slice, length + 1
	s = append(s, []int{1, 2, 3, 5, 6})
	fmt.Println(len(s))
}

// make matrix
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			pic[i][j] = uint8(i * j)
		}
	}
	return pic
}

// map
func WordCount(s string) map[string]int {
	count_map := make(map[string]int)
	for _, w := range strings.Fields(s) {
		count_map[w] += 1
	}
	return count_map
}

// function as return
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, (a + b)
		return a
	}
}

func main() {
	pointer()
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
