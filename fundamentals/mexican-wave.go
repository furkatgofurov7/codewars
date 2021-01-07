package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "hello"
	//var res []string 
	for i, _ := range word {
		var a []string
		newSlice := []string(append([]string(strings.ToUpper(string(word[i]))), word[i+1:])))
		a = append(a, newSlice)

		fmt.Println(a)
		
	}
	

}