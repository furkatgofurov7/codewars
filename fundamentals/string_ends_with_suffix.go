//https://www.codewars.com/kata/51f2d1cafc9c0f745c00037d

package main

import (
	"fmt"
	"strings"
	//"fmt"
	//"strconv"
	//"sort"
)

func main() {
	solution("Hello", "ll")
	solution("allo", "lo")
	solution("Hello", "l0")
}

func solution(str, ending string) bool {
	// Your code here!
	if strings.HasSuffix(str, ending) {
		fmt.Println("True! 1st argument ends with 2nd argument")
		return true
	}
	fmt.Println("False! 1st argument doesn't end with 2nd argument")
	return false 
}

// 2nd solution
// func solution(str, ending string) bool {
// 	return strings.HasSuffix(str, ending)
// }

// 3rd solution
// func solution(str, ending string) bool {
// 	return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
// }