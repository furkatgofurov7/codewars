//https://www.codewars.com/kata/56fa3c5ce4d45d2a52001b3c

package main

import "fmt"

func main() {
	Xor(true, false)
	Xor(false, true)
	Xor(true, true)
	Xor(false, false)
}

func Xor(a, b bool) bool {
	if a != b {
		fmt.Println("true")
		return true
	}
	fmt.Println("false")
	return false
}