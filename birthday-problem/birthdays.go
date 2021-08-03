package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const runs = 1_000_000

	for i := 2; i <= 23; i++ {
		fmt.Printf("P(two birthdays on same day for %d people) = %.2f%%\n", i, simulateBirthdays(i, runs)*100)
	}

	fmt.Printf("Results determined over %d runs", runs)
}

func simulateBirthdays(numberOfPeople, runs int) float64 {
	same := 0
	for i := 0; i < runs; i++ {
		if simulateBirthdayMatches(numberOfPeople) {
			same++
		}
	}

	return float64(same) / float64(runs)
}

func simulateBirthdayMatches(numberOfPeople int) bool {
	const daysInYear = 365

	seen := make(map[int]bool)
	for i := 0; i < numberOfPeople; i++ {
		day := rand.Intn(daysInYear)
		if seen[day] {
			return true
		}
		seen[day] = true
	}

	return false
}
