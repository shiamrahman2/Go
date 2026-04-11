package main

import "fmt"

func practice() {
	// print
	fmt.Println("Hello, World!")
	// variable declaration
	var var_name int = 10
	fmt.Println(var_name)
	var a = "Shiam Hosen Monna"
	fmt.Println(a)
	b := true
	fmt.Println(b)
	// reassign
	b = false
	fmt.Println(b)
	num := 0

	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
	c := 4
	switch c {
	case 1:
		fmt.Println("Number is one")
	case 2:
		fmt.Println("Number is Two")
	case 3:
		fmt.Println("Number is Three")
	case 4:
		fmt.Println("Number is Four")
	default:
		fmt.Println("Number Not Matched")
	}
}
func input() {
	fmt.Println("Welcome to our Application")
	var name string
	fmt.Println("Enter your name-")
	fmt.Scanln(&name)
	fmt.Println("Thanks ", name, " to use our Application")
}
func sum(num1 int, num2 int) {
	Sum := num1 + num2
	fmt.Println(Sum)
}

// function with return type
func ad(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

// function with multiple return value
func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}
func main() {
	// practice();
	// sum(10,20)
	//  a:=30
	//  b:=-3
	//  //value:=add(a,b)
	//  //fmt.Println(value)
	//  sum,mul:=getNumbers(a,b)
	//  fmt.Println(sum);
	//  fmt.Println(mul)
	//input()
	add(5, 7)
}
