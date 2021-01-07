package main

import "fmt"

func Summation(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
	  res += i
	}
	fmt.Println(res)
	return res
}

// func Summation(n int) (sum int) {
// 	//res := 0
// 	for i := 0; i <= n; i++ {
// 	  sum += i
// 	}
// 	fmt.Println(sum)
// 	//return res
// 	return
// }


// func Summation(n int) int {
// 	fmt.Println(n * (n + 1) / 2)
// 	return n * (n + 1) / 2
// }


// func Summation(n int) int {
//     if n == 1 {
//       return n
//     }
//     return n + Summation(n - 1)
// }

func main() {
	Summation(5)
}