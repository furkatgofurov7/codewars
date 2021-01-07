/*

Complete the method that takes a boolean value and return a "Yes"
string for true, or a "No" string for false.

*/

package main

// imported package(s)
// import (
// 	//"fmt"
// 	//"strings"
// 	//"fmt"
// 	//"strconv"
// 	//"sort"
// )

// BoolToWord returns yes or no
func BoolToWord(word bool) string {
	if word {
		//fmt.Println("Yes")
		return "Yes"
	}
	// fmt.Println("No")
	return "No"
}

// 2nd solution
// func BoolToWord(word bool) string {
// 	 return map[bool]string{false:"No", true:"Yes"}[word]
// 	//a := map[bool]string{false:"No", true:"Yes"}[word]
// 	//fmt.Println(a)
// 	//return a
// }

// 3rd solution
// func solution(str, ending string) bool {
// 	return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
// }


// For testing the solution
func main() {
	BoolToWord(true)
	BoolToWord(false)
}