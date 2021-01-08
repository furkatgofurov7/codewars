//https://www.codewars.com/kata/55cbc3586671f6aa070000fb

package main

func main() {
	CheckForFactor(7, 2)
}

func CheckForFactor(base int, factor int) bool {
  if base % factor == 0 {
    return true
  }
  return false
}