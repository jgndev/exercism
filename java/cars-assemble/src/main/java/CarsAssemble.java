public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        final int CARS_PER_HOUR = 221;
        final double LOW_ERROR_RATE = 0.9;
        final double MEDIUM_ERROR_RATE = 0.8;
        final double HIGH_ERROR_RATE = 0.77;

        return switch (speed) {
            case 1, 2, 3, 4 -> speed * CARS_PER_HOUR;
            case 5, 6, 7, 8 -> speed * CARS_PER_HOUR * LOW_ERROR_RATE;
            case 9 -> speed * CARS_PER_HOUR * MEDIUM_ERROR_RATE;
            case 10 -> speed * CARS_PER_HOUR * HIGH_ERROR_RATE;
            default -> 0.0;
        };
    }

    public int workingItemsPerMinute(int speed) {
        final int MINUTES_PER_HOUR = 60;

        var workingCarsPerHour = productionRatePerHour(speed);
        return (int)workingCarsPerHour / MINUTES_PER_HOUR;
    }
}
