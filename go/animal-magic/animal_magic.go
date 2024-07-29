package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	// Add one to the resulting random number because the possible
	// range is 0 to 19 with a provided argument of 20.
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	// rand.Float64 returns a random float64 between 0.0 and 1.0.
	// Multiply the result by 12.0 to make the range from 0.0 to 12.0
	return rand.Float64() * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	shuffled := make([]string, len(animals))
	copy(shuffled, animals)

	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
