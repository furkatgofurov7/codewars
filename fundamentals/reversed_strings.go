//https://www.codewars.com/kata/5168bb5dfe9a00b126000018

package main

// import "fmt"

func main() {
	Solution("hello")
}

func Solution(word string) string {
  reverse := []rune(word)
  for i, ch := 0, len(reverse)-1; i < ch; i, ch = i+1, ch-1 {
    reverse[i], reverse[ch] = reverse[ch], reverse[i]
  }
  return string(reverse)
}

/* 2nd and best solution:
func Solution(word string) string {
	var result = ""
	for _, c := range word {
	  result = string(c) + result
	  fmt.Println(result)
	}
	fmt.Println(result)
	return result
}
*/

