package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	ratePerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	ratePerMin := ratePerHour / 60
	return int(ratePerMin)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	batchCount := int(carsCount / 10)
	remainderCount := carsCount % 10

	batchCost := batchCount * 95000
	remainderCost := remainderCount * 10000

	return uint(batchCost + remainderCost)
}
