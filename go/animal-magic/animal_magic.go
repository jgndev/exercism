package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    // Add one so the range is between 1 and 20. rand.Intn is "half open", so it counts "up to" the provided argument
    // from 0 to n. We want the lowest value to be a 1, and the highest to be 20, which is achieved by adding 1.
    return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    // rand.Float64() returns a random float64 between 0.0 and 1.0. To get the desired range of 0.0 to 12.0
    // multiply the result by 12.
    return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

    // rand.Shuffle rearranges the values in the slice "in place".
    rand.Shuffle(len(animals), func(i, j int) {
        animals[i], animals[j] = animals[j], animals[i]
    })

    return animals
}
