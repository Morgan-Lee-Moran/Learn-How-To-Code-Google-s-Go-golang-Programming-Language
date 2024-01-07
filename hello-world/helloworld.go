package main

import (
	"fmt"
	"strconv"
	"strings"
)

func helloworld() {
	fmt.Println("Hello, World !")
}

func printf() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)
}

func unquoteCodePoint(s string) (string, error) {
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	return string(r), err
}

func printEmoji() {
	emoji, error := unquoteCodePoint("1f680")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf(emoji)
}

func basicType() {
	var a int = 10
	// define type when declear variable
	// a cannot be string or other type
	fmt.Println(a)
	var b = 12
	// b can not be any type either
	// init as string so it must be string
	fmt.Println(b)
	var s = "hello world"
	fmt.Println(s)
	var c = 'a'
	fmt.Println(c)
	var bo = true
	fmt.Println(bo)
	var f = 1.2
	fmt.Println(f)
	// fixed size array
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	// slice, dynamic size array
	var arr2 = [...]int{1, 2, 3}
	fmt.Println(arr2)
	var myMap = map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap)
	var map2 = make(map[string]int)
	fmt.Println(map2)
	var slice = []int{1, 2, 3}
	fmt.Println(slice)
	var slice2 = make([]int, 3)
	fmt.Println(slice2)
	// init with 3 element ane have capacity of 5
	var slice3 = make([]int, 3, 5)
	fmt.Println(slice3)
	var slice4 = arr[0:2]
	fmt.Println(slice4)
	var slice5 = arr[1:]
	fmt.Println(slice5)
	var slice6 = arr[:2]
	fmt.Println(slice6)
	var slice7 = append(slice, 4)
	fmt.Println(slice7)
	var slice8 = append(slice, slice2...)
	fmt.Println(slice8)

	// name is a type and it is same as string
	type name string

	// convert int to float64
	var i int = 42
	var f2 float64 = float64(i)
	fmt.Println(f2)

	// interface{} is any type
	var ix = interface{}(42)
	fmt.Println(ix)

	//get type by interface{}
	var is interface{} = "hello"
	si := is.(string)
	fmt.Println(si)
}

func main() {
	helloworld()
	printf()
	printEmoji()
	basicType()
}
