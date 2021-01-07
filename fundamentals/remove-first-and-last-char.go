/*

Problem Description:

It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. 
You're given one parameter, the original string. You don't have to worry with strings with less than two characters.

*/

package main

//import "fmt"

func RemoveChar(word string) string {
	//fmt.Println(word[1:len(word)-1])
	return word[1:len(word)-1]
}

func main() {
	RemoveChar("Hello World!")
	RemoveChar("Hello               hey!s")
}