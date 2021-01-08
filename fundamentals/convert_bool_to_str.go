//https://www.codewars.com/kata/53369039d7ab3ac506000467

package main

func main() {
	BoolToWord(true)
	BoolToWord(false)
}

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	}
	return "No"
}

// 2nd solution
// func BoolToWord(word bool) string {
// 	return map[bool]string{false:"No", true:"Yes"}[word]
// 	a := map[bool]string{false:"No", true:"Yes"}[word]
// 	return a
// }

// 3rd solution
// func solution(str, ending string) bool {
// 	return len(str) >= len(ending) && str[len(str) - len(ending):] == ending
// }