package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car" {
        return true
    } else if kind == "bike" {
    	return false
    } else if kind == "truck" {
		return true
    } else {
    	return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
        return fmt.Sprintf("%s is clearly the better choice.", option1)
    } else {
        return fmt.Sprintf("%s is clearly the better choice.", option2)
    }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	depreciation := 0.0
    if age < 3.0 {
        depreciation = 0.8
    } else if (age >= 3.0 && age < 10.0) {
    	depreciation = 0.7
    } else {
    	depreciation = 0.5
    }
	return originalPrice * depreciation
}
