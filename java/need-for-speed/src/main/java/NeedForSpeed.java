class NeedForSpeed {
    // Normal car instance variables
    private int speed;
    private int batteryDrain;
    private int batteryLevel = 100;
    private int distanceDriven;

    // Special Nitro Edition speed and battery drain
    private static final int NITRO_SPEED = 50;
    private static final int NITRO_BATTERY_DRAIN = 4;

    NeedForSpeed(int speed, int batteryDrain) {
        if (speed <= 0 || batteryDrain <= 0) {
            throw new IllegalArgumentException("Speed and Battery Drain must be a positive integer.");
        }

        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public boolean batteryDrained() {
        return this.batteryLevel <= 0;
    }

    public int distanceDriven() {
        return this.distanceDriven;
    }

    public void drive() {
        if (!batteryDrained()) {
            this.distanceDriven += this.speed;
            this.batteryLevel -= this.batteryDrain;
        }
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(NITRO_SPEED, NITRO_BATTERY_DRAIN);
    }

    public int getSpeed() {
        return this.speed;
    }

    public int getBatteryDrain() {
        return this.batteryDrain;
    }

    public int getBatteryLevel() {
        return this.batteryLevel;
    }
}

class RaceTrack {
    private int trackDistance = 0;

    RaceTrack(int distance) {
        this.trackDistance = distance;
    }

    public boolean tryFinishTrack(NeedForSpeed car) {
        var drivesNeeded = (getTrackDistance() + car.getSpeed() - 1) / car.getSpeed();
        var batteryNeeded = car.getBatteryDrain() * drivesNeeded;

        return batteryNeeded <= car.getBatteryLevel();
    }

    // Getters and Setters for private fields
    public int getTrackDistance() {
        return trackDistance;
    }
}
