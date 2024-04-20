package elon

import "fmt"

// Drive increases the distance driven by the Car's speed while decreasing
// the amount of available battery by the Car's battery drain.
func (c *Car) Drive() {
    // Car's can only drive if they have available battery to use.
    if c.battery-c.batteryDrain >= 0 {
        c.distance += c.speed
        c.battery -= c.batteryDrain
    }
}

// DisplayDistance returns a string of the format "Driven n meters", where
// n is the distance driven so far from one or more calls to Drive().
func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns a string of the format "Battery at n%", where
// n is the amount of battery the car currently has.
func (c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish returns a bool that represents whether or not the Car can finish
// a race with a track of n length.
func (c *Car) CanFinish(trackDistance int) bool {
    // Ceiling Division, ensure that any remainder in the division will push
    // the quotient up to the next integer. Useful for when remainders discarded
    // by integer division would be thrown away and give an incorrect result.
    drivesNeeded := (trackDistance + c.speed - 1) / c.speed
    batteryNeeded := drivesNeeded * c.batteryDrain

    return batteryNeeded <= c.battery
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
