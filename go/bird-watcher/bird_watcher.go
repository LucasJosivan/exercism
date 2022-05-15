package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (result int) {
	for _, birds := range birdsPerDay {
		result += birds
	}
	return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	offset := 7 * (week - 1)
	return TotalBirdCount(birdsPerDay[offset : offset+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for n := range birdsPerDay {
		if n%2 == 0 {
			birdsPerDay[n]++
		}
	}
	return birdsPerDay
}
