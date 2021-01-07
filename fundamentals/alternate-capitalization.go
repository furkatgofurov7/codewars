package main

import (
	//"fmt"
	"fmt"
	//"unicode"
	"strings"
)

// func Capitalize(s string) []string {
// 	a, b := []rune(s),[]rune(s)
// 	for i := range a {
// 	  if i%2 == 0 {
// 		a[i] = unicode.ToUpper(a[i])
// 	  }else{
// 		b[i] = unicode.ToUpper(b[i])
// 	  }
// 	}
// 	fmt.Println(string(a), string(b))
// 	return []string{string(a), string(b)}
// }

func Capitalize(st string) []string {
	s1 := ""
	s2 := ""
	for i, c :=range st {
	  if i%2==0 {
		s1 += strings.ToUpper(string(c))
		s2 += strings.ToLower(string(c))
	  } else {
		s1 += strings.ToLower(string(c))
		s2 += strings.ToUpper(string(c))
	  }
	}
	fmt.Println(string(s1), string(s2))
	return []string{s1,s2}
  }

// func Capitalize(st string) []string {
// 	stringEven := make([]string, len(st))
// 	stringOdd := make([]string, len(st))
	
	
// 	for i := 0; i < len(st); i++ { 
// 		if i % 2 == 0 {
// 		stringEven[i] = strings.ToUpper(string(st[i]))
// 	  } else {
// 		stringEven[i] = string(st[i])
// 	  }
// 	}

// 	for i := 0; i < len(st); i++ { 
// 		if i % 2 != 0 {
// 		stringOdd[i] = strings.ToUpper(string(st[i]))
// 	  } else {
// 		stringOdd[i] = string(st[i])
// 	  }
// 	}
// 	reg := []string {string(stringOdd), stringEven}
// 	//return (stringEven[i], stringOdd[i])
// 	fmt.Println(stringEven,stringOdd)
// 	return stringEven
// }

func main() {
	Capitalize("abcdefgh")
}