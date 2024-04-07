public class CarsAssemble {

    private static final int CARS_PER_HOUR = 221;
    private static final int MINUTES_PER_HOUR = 60;
    private static final double LOW_ERROR_RATE = 0.9;
    private static final double MED_ERROR_RATE = 0.8;
    private static final double HIGH_ERROR_RATE = 0.77;

    public double productionRatePerHour(int speed) {
        return switch (speed) {
            case 1, 2, 3, 4 -> CARS_PER_HOUR * speed;
            case 5, 6, 7, 8 -> CARS_PER_HOUR * speed * LOW_ERROR_RATE;
            case 9 -> CARS_PER_HOUR * speed * MED_ERROR_RATE;
            case 10 -> CARS_PER_HOUR * speed * HIGH_ERROR_RATE;
            default -> 0;
        };
    }

    public int workingItemsPerMinute(int speed) {
        return (int)productionRatePerHour(speed) / MINUTES_PER_HOUR;
    }
}
