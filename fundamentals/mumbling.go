/*
Problem Description:

This time no story, no theory. The examples below show you how to write function accum:

Examples:

accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"

The parameter of accum is a string which includes only letters from a..z and A..Z.
*/

package main

import (
	"fmt"
	"strings"
	//"unicode"
)

func Accum(s string) string {
	parts := make([]string, len(s))
	//fmt.Println(parts)
    
    for i := 0; i < len(s); i++ {
      parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
    }
    fmt.Println(strings.Join(parts, "-"))
    return strings.Join(parts, "-")
}


//solution2
/*
func Accum(s string) string {
    words := make([]string, len(s))
    
    for i, c := range s {
        words[i] = strings.Title(strings.Repeat(strings.ToLower(string(c)), i+1))
    }
   
    return strings.Join(words, "-")
}
*/

//solution3
/*
func Accum(s string) string {
    res := ""
    for i, r := range s {
        if i != 0 {
            res += "-"
        }
        res += string(unicode.ToTitle(r))
        for j := 0; j < i; j++ {
            res += string(unicode.ToLower(r))
        }
    }
    return res
}
*/

//solution4
/*
func Accum(s string) string {
    out := ""
    i := 1;
    for _, r := range s {
        c := string(r)
      
      out += strings.ToUpper(c)
      out += strings.Repeat(strings.ToLower(c),i-1)
      out += "-"
      i++
    }
    return out[0:len(out)-1]
}
*/

//solution5
/*
func Accum(s string) string {
	var result string
	for pos, char := range s {
	  c := strings.ToLower(string(char))
	  segment := strings.Repeat(c, pos+1)
	  result += strings.Title(segment) + "-"
	}
	return result[:len(result)-1] // slice off final "-"
  }
*/


func main() {
	Accum("abcd")
	Accum("RqaEzty")
	Accum("RRddwdag")
	Accum("cwAt")
	Accum("pmBeW")
}