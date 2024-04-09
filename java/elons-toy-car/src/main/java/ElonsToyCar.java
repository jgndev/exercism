public class ElonsToyCar {

    private float batteryCharge = 100;
    private float distanceDriven = 0;

    /**
     * Factory method to purchase a new Elon's Toy Car.
     *
     * @return A new instance of {@code ElonsToyCar}.
     */
    public static ElonsToyCar buy() {
        return new ElonsToyCar();
    }

    /**
     * Returns a display string indicating the total distance driven.
     *
     * @return A string in the format "Driven {distance} meters", where {distance} is the total distance driven by the car.
     */
    public String distanceDisplay() {
        return String.format("Driven %.0f meters", this.distanceDriven);
    }

    /**
     * Returns a display string indicating the current battery charge.
     * If the battery charge is above 0%, the returned string is "Battery at {charge}%",
     * where {charge} is the current battery charge. If the battery charge is 0%, it returns "Battery empty".
     *
     * @return A string representing the current battery status.
     */
    public String batteryDisplay() {
        if (this.batteryCharge > 0) {
           return String.format("Battery at %.0f%%", this.batteryCharge);
        } else {
            return "Battery empty";
        }
    }

    /**
     * Simulates driving the toy car, reducing the battery charge by 1% and increasing the distance driven by 20 meters.
     * If the battery is empty, the car cannot drive any further.
     */
    public void drive() {
        if (this.batteryCharge > 0) {
            this.distanceDriven += 20.0f;
            this.batteryCharge -= 1.0f;
            batteryDisplay();
        }
    }

}
