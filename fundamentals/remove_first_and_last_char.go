//https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0

package main

func main() {
	RemoveChar("Hello World!")
	RemoveChar("Hello               hey!s")
}

func RemoveChar(word string) string {
	return word[1:len(word)-1]
}