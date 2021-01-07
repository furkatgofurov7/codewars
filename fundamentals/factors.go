/*

Problem Description:

This function should test if the factor is a factor of base.

Return true if it is a factor or false if it is not.

About factors

Factors are numbers you can multiply together to get another number.

2 and 3 are factors of 6 because: 2 * 3 = 6

    You can find a factor by dividing numbers. If the remainder is 0 then the number is a factor.
    You can use the mod operator (%) in most languages to check for a remainder

For example 2 is not a factor of 7 because: 7 % 2 = 1

Note: base is a non-negative number, factor is a positive number.
*/

package main

import "fmt"

func CheckForFactor(base int, factor int) bool {
	if base % factor == 0 {
		fmt.Println("It is a factor")
		return true
	}
	fmt.Println("It is not a factor")
	return false
}

func main() {
	CheckForFactor(9, 2)
	CheckForFactor(5412, 6)
	CheckForFactor(192, 2)
	CheckForFactor(653, 7)
}