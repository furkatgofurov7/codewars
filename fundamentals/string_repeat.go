//https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

package main

import (
	"strings"
)

func main() {
	RepeatStr(6, "Hello")
}

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

// 2nd solution
// func RepeatStr(repetitions int, value string) (res string) {
// 	for i := 0; i < repetitions; i++ {
// 		res += value
// 	}
// 	fmt.Println(res)
// 	return
// }