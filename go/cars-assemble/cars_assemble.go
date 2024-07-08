package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    sr := successRate / 100
    pr := float64(productionRate)
    return pr * sr
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    cph := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(cph) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    tens := carsCount / 10
    tensCost := tens * 9_5000

    singles := carsCount % 10
    singlesCost := singles * 10_000

    return uint(tensCost + singlesCost)
}
