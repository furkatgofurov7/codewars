/*

Problem Description:

Grade book

Complete the function so that it finds the mean of the three scores passed to it and returns the letter value associated with that grade.
Numerical Score 	Letter Grade
90 <= score <= 100 	'A'
80 <= score < 90 	'B'
70 <= score < 80 	'C'
60 <= score < 70 	'D'
0 <= score < 60 	'F'

Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.
*/

package main

import "fmt"

func GetGrade(a,b,c int) rune {
    // Code here
  var r1, r2, r3, r4, r5 rune
  r1 = 'A'
  r2 = 'B'
  r3 = 'C'
  r4 = 'D'
  r5 = 'F'
  res := (a+b+c)/3
  if res >= 90 && res <= 100 {
    fmt.Println(string(r1))
    return r1
  } else if res >= 80 && res < 90 {
    fmt.Println(string(r2))
    return r2
  } else if res >= 70 && res < 80 {
    fmt.Println(string(r3))
    return r3
  } else if res >= 60 && res < 70 {
    fmt.Println(string(r4))
    return r4
  } else {
    fmt.Println(string(r5))
    return r5
  }
}

/* 2nd solution:
func GetGrade(a, b, c int) rune {
    switch ((a + b + c) / 30) {
        case 10: return 'A'
        case 9: return 'A'
        case 8: return 'B'
        case 7: return 'C'
        case 6: return 'D'
        default: return 'F'
    }
}

3rd solution:
func GetGrade(a,b,c int) rune {
    switch mean := (a+b+c)/3; {
    case mean < 60: 
    return 'F'
    case mean < 70: 
    return 'D'
    case mean < 80: 
    return 'C'
    case mean < 90: 
    return 'B'
    default: 
    return 'A'
    }
    
}

4th solution:
func GetGrade(a,b,c int) rune {
    var sum = (a + b + c) / 3
    
    if sum >= 90 {
      return 'A'
    } else if sum >= 80 {
      return 'B'
    } else if sum >= 70 {
      return 'C'
    } else if sum >= 60 {
      return 'D'
    } else {
      return 'F'
    }
}*/

func main() {
  GetGrade(95, 95, 90)
  GetGrade(55, 75, 60) 
  GetGrade(55, 65, 50) 
  GetGrade(85, 85, 70)
}