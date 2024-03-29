public class Lasagna {
    private final int EXPECTED_TIME_IN_OVEN = 40;
    private final int MINUTES_PER_LAYER = 2;

    public int expectedMinutesInOven() {
        return EXPECTED_TIME_IN_OVEN;
    }

    public int remainingMinutesInOven(int elapsedMinutes) {
        return expectedMinutesInOven() - elapsedMinutes;
    }

    public int preparationTimeInMinutes(int layers) {
        return MINUTES_PER_LAYER * layers;
    }

    public int totalTimeInMinutes(int layers, int ovenTime) {
        return preparationTimeInMinutes(layers) + ovenTime;
    }
}
