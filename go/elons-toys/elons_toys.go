package elon

import "fmt"

// Drive updates the number of meters driven based on the car's speed,
// and reduces the battery according to the battery drain.
func (c *Car) Drive() {
	if c.battery-c.batteryDrain >= 0 {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// DisplayDistance returns a string message of the distance driven.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns a string message of the remaining battery.
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish returns a bool result of determining if the car can complete
// a track distance of 'trackDistance' meters.
func (c *Car) CanFinish(trackDistance int) bool {
	possibleDrives := c.battery / c.batteryDrain
	possibleDistance := possibleDrives * c.speed
	return possibleDistance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
