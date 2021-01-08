//https://www.codewars.com/kata/55f73be6e12baaa5900000d4

package main

func main() {
	Goals(300, 45, 100)
}

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return (laLigaGoals+copaDelReyGoals+championsLeagueGoals)
}