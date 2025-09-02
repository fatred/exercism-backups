package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for day := 0; day < len(birdsPerDay); day++ {
        sum = sum + birdsPerDay[day]
    }
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    endingOffset := (week * 7)
	return TotalBirdCount(birdsPerDay[endingOffset-7 : endingOffset])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day := 0; day < len(birdsPerDay) ; day++ {
    	if day % 2 == 0 {
        	birdsPerDay[day] = birdsPerDay[day] + 1    
        }
    }
	return birdsPerDay
}
