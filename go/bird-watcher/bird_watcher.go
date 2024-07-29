package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdCount := 0

	for i := 0; i < len(birdsPerDay); i++ {
		birdCount += birdsPerDay[i]
	}

	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsInWeek := 0

	// week 1 = 0    1 - 1 = 0, 0 * 7 = 0
	// week 2 = 7    2 - 1 = 1, 1 * 7 = 7
	// week 3 = 14   3 - 1 = 2, 2 * 7 = 14
	// week 4 = 21   4 - 1 = 3, 3 * 7 = 21
	startIndex := (week - 1) * 7

	for i := startIndex; i <= startIndex+7 && i < len(birdsPerDay); i++ {
		birdsInWeek += birdsPerDay[i]
	}

	return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// Increase bird count day by one on odd days starting from index 0
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}

	return birdsPerDay
}
