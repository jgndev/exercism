package speed

// Car is a struct that represents the properties of a toy car
type Car struct {
    battery      int
    batteryDrain int
    speed        int
    distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    return Car{
        battery:      100,
        batteryDrain: batteryDrain,
        speed:        speed,
        distance:     0,
    }
}

// Track is a struct that represents a track for testing toy cars.
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
    return Track{
        distance: distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery >= car.batteryDrain {
        car.battery -= car.batteryDrain
        car.distance += car.speed
    }

    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    // How many drives will it take to complete the track at car.speed
    // Account for track distance that should have a remainder when dividing by car.speed
    drivesNeeded := (track.distance + car.speed - 1) / car.speed

    // How much battery is needed to complete the required drives
    batteryNeeded := car.batteryDrain * drivesNeeded

    // Is the battery needed less than the available battery
    return batteryNeeded <= car.battery
}
