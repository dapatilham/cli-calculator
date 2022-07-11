package main

import (
	"flag"
	"fmt"
)

func add(first int, second int) (result int) {
	result = first + second
	return
}

func sub(first int, second int) (result int) {
	result = first - second
	return
}

func mul(first int, second int) (result int) {
	result = first * second
	return
}

func div(first int, second int) (result int) {
	result = first / second
	return
}

func mod(first int, second int) (result int) {
	result = first % second
	return
}

func main() {

	var first, second int
	var op string

	flag.IntVar(&first, "first", 0, "-first= ")
	flag.StringVar(&op, "op", "", "-op= ")
	flag.IntVar(&second, "second", 0, "-second= ")

	flag.Parse()

	switch {
	case op == "+":
		total := add(first, second)
		fmt.Printf("Result of %d + %d = %d \n", first, second, total)
	case op == "-":
		total := sub(first, second)
		fmt.Printf("Result of %d - %d = %d \n", first, second, total)
	case op == "*":
		total := mul(first, second)
		fmt.Printf("Result of %d * %d = %d \n", first, second, total)
	case op == "/":
		total := div(first, second)
		fmt.Printf("Result of %d / %d = %d \n", first, second, total)
	case op == "%":
		total := mod(first, second)
		fmt.Printf("Result of %d % %d = %d \n", first, second, total)
	default:
		fmt.Println("Invalid operation")
	}

}