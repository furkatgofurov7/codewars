//https://www.codewars.com/kata/5667e8f4e3f572a8f2000039

package main

import (
	"strings"
)

func main() {
	Accum("abcd")
	Accum("RqaEzty")
	Accum("RRddwdag")
	Accum("cwAt")
	Accum("pmBeW")
}

func Accum(s string) string {
	parts := make([]string, len(s))
    for i := 0; i < len(s); i++ {
      parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
    }
    return strings.Join(parts, "-")
}

/* 2nd solution
func Accum(s string) string {
    words := make([]string, len(s))
    
    for i, c := range s {
        words[i] = strings.Title(strings.Repeat(strings.ToLower(string(c)), i+1))
    }
   
    return strings.Join(words, "-")
}

3rd solution
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

4th solution:
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

5th solution:
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