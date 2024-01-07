package main

import "fmt"

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

// example of define a struct
type HelloWorld struct {
	Name string
	Age  int
}

// constructor
func newHelloWorldNameAge(name string, age int) HelloWorld {
	return HelloWorld{
		Name: name,
		Age:  age,
	}
}

// constructor overloading is not supported in golang
func newHelloWorldName(name string) HelloWorld {
	return HelloWorld{
		Name: name,
		Age:  18,
	}
}

func printHelloWorld(h HelloWorld) {
	// go do not support direct concat of string and int
	// "Hellow " + h.age is invalid
	fmt.Printf("Hello %s age: %d\n", h.Name, h.Age)
}

func callHelloWorldConstructorAndPrint() {
	var hello HelloWorld = newHelloWorldNameAge("Kim", 22)
	printHelloWorld(hello)
}

// example of declear a interface and implement
/**
In Go, unlike languages like Java, there is no explicit declaration of intent when a type implements an interface. Instead, Go uses a concept known as structural typing or duck typing for interfaces. This means that if a type has all the methods that an interface requires, it implicitly implements that interface. There's no need to explicitly state that a type implements a given interface.
*/
type toStringInterface interface {
	toString() string
}

func (h HelloWorld) toString() string {
	return fmt.Sprintf("Hi %s age: %d\n", h.Name, h.Age)
}

func callInterfaceToString(t toStringInterface) {
	fmt.Println(t.toString())
}

func callHelloWorldConstructorAndInterface() {
	var hello HelloWorld = newHelloWorldName("Lim")
	callInterfaceToString(hello)
}

func main() {
	basicType()
	callHelloWorldConstructorAndPrint()
	callHelloWorldConstructorAndInterface()
}
