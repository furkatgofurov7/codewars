/*

Problem Description:

Complete the solution so that it returns true if the first argument(string) passed
in ends with the 2nd argument (also a string).

Examples:

solution("abc", "bc") // returns true
solution("abc", "d") // returns false

*/

package main

// imported package(s)
import (
	"fmt"
	"strings"
	//"fmt"
	//"strconv"
	//"sort"
)

//solution returns true or false
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


// For testing the solution
func main() {
	solution("Hello", "ll")
	solution("allo", "lo")
	solution("Hello", "l0")
}