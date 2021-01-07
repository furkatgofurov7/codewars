package main

import (
	"fmt"
	"strings"
	//"math"
)


func High(s string) string {
	//letters := map[string]int
	letters := map[string]int {
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	//fmt.Println(letters["m"])
	split := strings.Split(s, " ")
	fmt.Println(split)
	fmt.Println(len(split))
	//fmt.Println(len(split))
	

	//var score_list [len(split)]int
	var score_list []int
	for i, word := range split {
		fmt.Printf("index is %v, word is %v\n", i, word)
		//word_score := make(map[string]int)
		
		//getScore = map[string(word)]int(word_score)
		//var score_list [split]int
		getScore := make(map[string]int)
		//fmt.Println(getScore)
		var word_score int
		
		for _, char := range word {
			//word_score := make(map[string]int)
			
			//word_score[j] = letters[string(char)]
			//fmt.Printf("index is %v, character is %v\n", j, string(char))
			//word_score[string(char)] += letters[string(char)]
			word_score += letters[string(char)]
			//fmt.Println(word_score)
		}
		getScore = map[string]int{word:word_score}
		
		//fmt.Println(getScore)
		//fmt.Println(getScore[word])
		
		score_list = append(score_list, getScore[word])
		//fmt.Println(score_list)
		var max = score_list[0]
		for i := range score_list {
			if score_list[i] > max {
				max = score_list[i]
				//word_score = max
				//fmt.Println(max)
				for key, value := range getScore {
					if value == max {
						
						//return key
					}
					//fmt.Println(key)
					return key
				}
				
				//fmt.Println(getScore[word])
			}
		}
		//fmt.Println(score_list)
		//fmt.Println(math.Max(score_list))
		//fmt.Println(strings.Join(getScore[word], "-"))
		// for j, char := range word {
		// 	word_score[j] := letters[char] + word_score[j]
			
		// }
		 
	}
	// score_list = append(score_list)
	// fmt.Println(score_list)
	// getScore := make(map[string]int)
	// getScore[]
	// fmt.Println(strings.Join(getScore, ","))
	return strings.Join(split, "-")
}



func main() {
	High("my name is Shakhnoza shakhnoza sas asasas sasa sassaaa")
}