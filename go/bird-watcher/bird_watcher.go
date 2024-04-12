package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    birdCount := 0

    for _, count := range birdsPerDay {
        birdCount += count
    }

    return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    // []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
    //       0  1  2  3  4  5  6  0  1  2  3  4  5  6
    start := (week - 1) * 7
    end := start + 7
    // Bounds check to make sure end isn't longer than birdsPerDay
    if end > len(birdsPerDay) {
        end = len(birdsPerDay)
    }

    // Take a slice from start to end
    birds := birdsPerDay[start:end]

    birdCount := 0

    for _, count := range birds {
        birdCount += count
    }

    return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days. Modifies the slice in place.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] += 1
    }

    return birdsPerDay
}
