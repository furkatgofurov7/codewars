package main

import (
	"strings"
	"fmt"
)


// func High(s string) string {
// 	letters := map[string]int {
// 		"a": 1,
// 		"b": 2,
// 		"c": 3,
// 		"d": 4,
// 		"e": 5,
// 		"f": 6,
// 		"g": 7,
// 		"h": 8,
// 		"i": 9,
// 		"j": 10,
// 		"k": 11,
// 		"l": 12,
// 		"m": 13,
// 		"n": 14,
// 		"o": 15,
// 		"p": 16,
// 		"q": 17,
// 		"r": 18,
// 		"s": 19,
// 		"t": 20,
// 		"u": 21,
// 		"v": 22,
// 		"w": 23,
// 		"x": 24,
// 		"y": 25,
// 		"z": 26,
// 	}
// 	split := strings.Split(s, " ")
// 	var score_list []int
// 	for _, word := range split {
// 		getScore := make(map[string]int)
// 		var word_score int
		
// 		for _, char := range word {
// 			word_score += letters[string(char)]
// 		}
// 		getScore = map[string]int{word:word_score}	
// 		score_list = append(score_list, getScore[word])
// 		var max = score_list[0]
// 		for i := range score_list {
// 			if score_list[i] > max {
// 				max = score_list[i]
// 				for key, value := range getScore {
// 					if value == max {
// 					}
// 					return key
// 				}
// 			}
// 		}
// 	}
// 	return strings.Join(split, "-")
// }
func High(s string) string {
	split := strings.Split(s, " ")
	fmt.Println(split)
	for i, word := range split {
		fmt.Printf("index is %v, word is %v\n", i, word)
	}
	return fmt.Sprintf("%v", split)
}
func main() {
	High("my name is Shakhnoza aas jdad dhhhhh ssss")
}