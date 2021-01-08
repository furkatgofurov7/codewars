//https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a

package main

func main() {
	Past(1, 1, 1)
	Past(0, 0, 0) 
	Past(0, 1, 0) 
	Past(1, 0, 1)
}

func Past(h, m, s int) int {
	var hourInMs, minuteInMs, secondInMs int
	hourInMs = h * 3600000 
	minuteInMs = m * 60000
	secondInMs = s * 1000
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
}
*/