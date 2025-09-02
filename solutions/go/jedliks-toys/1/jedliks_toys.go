package jedlik

import "fmt"

func (car *Car) Drive() {
    if car.battery > car.batteryDrain {
        car.battery -= car.batteryDrain
        car.distance += car.speed
    }
}

func (car *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
    var canFinish bool = false
    maxDistance := (car.battery / car.batteryDrain) * car.speed
    if trackDistance <= maxDistance {
        canFinish = true
    }
    return canFinish
}

