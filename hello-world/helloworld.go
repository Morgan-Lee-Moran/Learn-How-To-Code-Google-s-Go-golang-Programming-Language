package main

import "fmt"

func helloworld() {
	fmt.Println("Hello, World !")
}

func printf() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)
}

func main() {
	helloworld()
	printf()
}
