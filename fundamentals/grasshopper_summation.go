//https://www.codewars.com/kata/55d24f55d7dd296eb9000030

package main

func main() {
	Summation(5)
}

func Summation(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
	  res += i
	}
	return res
}

// 2nd solution
// func Summation(n int) (sum int) {
// 	//res := 0
// 	for i := 0; i <= n; i++ {
// 	  sum += i
// 	}
// 	fmt.Println(sum)
// 	//return res
// 	return
// }

// 3rd solution
// func Summation(n int) int {
// 	fmt.Println(n * (n + 1) / 2)
// 	return n * (n + 1) / 2
// }

// 4th solution
// func Summation(n int) int {
//     if n == 1 {
//       return n
//     }
//     return n + Summation(n - 1)
// }