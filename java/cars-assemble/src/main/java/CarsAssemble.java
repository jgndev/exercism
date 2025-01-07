public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        return switch (speed) {
            case 1, 2, 3, 4 -> speed * 221;
            case 5, 6, 7, 8 -> speed * 221 * 0.9;
            case 9 -> speed * 221 * 0.8;
            case 10 -> speed * 221 * 0.77;
            default -> 0;
        };

    }

    public int workingItemsPerMinute(int speed) {
        return (int)productionRatePerHour(speed) / 60;
    }
}
