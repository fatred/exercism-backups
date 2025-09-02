// package Cars offers calculations relating to the effectiveness of a production line
package cars

import "math"

const ProductionCostPerCar int = 10000
const ProductionCostPerTenCars int = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRatePct := successRate / 100
    return float64(productionRate) * successRatePct
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
	remainder := math.Mod(float64(carsCount), 10.0)
	tensCost := ProductionCostPerTenCars * tens
	remainderCost := float64(ProductionCostPerCar) * remainder
	return uint(tensCost + int(remainderCost))}
