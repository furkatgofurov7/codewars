//https://www.codewars.com/kata/55cbd4ba903825f7970000f5

package main

func main() {
  GetGrade(95, 95, 90)
  GetGrade(55, 75, 60) 
  GetGrade(55, 65, 50) 
  GetGrade(85, 85, 70)
}

func GetGrade(a,b,c int) rune {
  var r1, r2, r3, r4, r5 rune
  r1 = 'A'
  r2 = 'B'
  r3 = 'C'
  r4 = 'D'
  r5 = 'F'
  res := (a+b+c)/3
  if res >= 90 && res <= 100 {
    return r1
  } else if res >= 80 && res < 90 {
    return r2
  } else if res >= 70 && res < 80 {
    return r3
  } else if res >= 60 && res < 70 {
    return r4
  } else {
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
}
*/