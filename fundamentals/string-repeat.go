/*

Problem Description:

Write a function called RepeatStr which repeats the given string value exactly repetitions count times.

repeatStr(6, "I") // "IIIIII"
repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"

*/
package main

// imported package(s)
import (
	"strings"
	"fmt"
	//"strconv"
	//"sort"
)

//RepeatStr returns string with N times repeated
func RepeatStr(repetitions int, value string) string {
	fmt.Println(strings.Repeat(value, repetitions))
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


// For testing the solution
func main() {
	RepeatStr(6, "Hello")
}