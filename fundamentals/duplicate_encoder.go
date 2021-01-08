/*

Problem Description:

The goal of this exercise is to convert a string to a new string 
where each character in the new string is "(" if that character 
appears only once in the original string, or ")" if that character 
appears more than once in the original string. Ignore capitalization
when determining if a character is a duplicate.

Examples

"din"      =>  "((("
"recede"   =>  "()()()"
"Success"  =>  ")())())"
"(( @"     =>  "))((" 

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
func DuplicateEncode(word string) string {
	var userNames []string  
	words := strings.Fields(word)
	fmt.Println(words)
	for _, char := range words {
		if strings.Count(word, char) > 1 {
			char = ")"
			//userNames = append(userNames, char)
		} else {
			char = "("
			//userNames = append(userNames, char)
		}		
		fmt.Printf("%v ", char)
	}
	return word
}

// For testing the solution
func main() {
	DuplicateEncode("din")
}