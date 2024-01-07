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

func main() {
	helloworld()
	printf()
	printEmoji()
}
