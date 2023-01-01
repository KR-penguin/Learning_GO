package myfunctions

import "fmt"

/*
func Calculate(a, b int) (plus, minus, multiplication, division int) {

	defer fmt.Println("Calculate Done")

	plus = a + b
	minus = a - b
	multiplication = a * b
	division = a / b
	return
}
*/

func Calculate(numbers ...int) (plus, minus, multiplication, division int) {

	defer fmt.Println("Calculate Done")

	for _, number := range numbers {
		plus += number
	}
	for _, number := range numbers {
		minus -= number
	}
	for _, number := range numbers {
		multiplication *= number
	}
	for _, number := range numbers {
		division /= number
	}

	return
}

//  func name(argument type) type { ... }
// go는 한 개의 function에서 1개이상의 값을 반환할 수 있다.
// 따라서 1개이상의 반환값을 가지는 함수를 선언할 때는
// func name(argument type) (type, type) { ... } 으로 선언하면 된다.

func multiply(a int, b int) int { // or func multiply(a, b int) (int) {
	return a * b
}