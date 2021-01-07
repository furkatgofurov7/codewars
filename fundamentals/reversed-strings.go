/*

Complete the solution so that it reverses the string passed into it.

'world'  =>  'dlrow'

*/

package main

import "fmt"

// func Solution(word string) string {
//   reverse := []rune(word)
//   for i, ch := 0, len(reverse)-1; i < ch; i, ch = i+1, ch-1 {
//     reverse[i], reverse[ch] = reverse[ch], reverse[i]
//   }
//   fmt.Println(string(reverse) + "\n")
//   return string(reverse)
// }

// 2nd and best solution
func Solution(word string) string {
	var result = ""
	for _, c := range word {
	  result = string(c) + result
	  fmt.Println(result)
	}
	fmt.Println(result)
	return result
}

func main() {
	Solution("hello")
	// Solution("RqaEzty")
	// Solution("RRddwdag")
	// Solution("cwAt")
	// Solution("pmBeW")
}