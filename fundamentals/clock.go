/*
Problem Description:

Clock shows 'h' hours, 'm' minutes and 's' seconds after midnight.

Your task is to make 'Past' function which returns time converted to milliseconds.
Example:

past(0, 1, 1) == 61000

Input constraints: 0 <= h <= 23, 0 <= m <= 59, 0 <= s <= 59
*/

package main

import "fmt"

func Past(h, m, s int) int {
	var hourInMs, minuteInMs, secondInMs int
	hourInMs = h * 3600000 
	minuteInMs = m * 60000
	secondInMs = s * 1000
	fmt.Println(hourInMs + minuteInMs + secondInMs)
	return (hourInMs + minuteInMs + secondInMs)
}

/* 2nd solution:
func Past(h, m, s int) int {
    return (h*3600000 + m*60000 + s*1000)
}

3rd solution:
func Past(h, m, s int) int {
    return (h*60*60+m*60+s)*1000
}

4th solution:
func Past(h, m, s int) int {
    milliseconds := s
    milliseconds += m * 60
    milliseconds += h * 3600
    milliseconds *= 1000

    return milliseconds
}*/

func main() {
	Past(1, 1, 1)
	Past(0, 0, 0) 
	Past(0, 1, 0) 
	Past(1, 0, 1)
}