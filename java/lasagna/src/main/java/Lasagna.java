public class Lasagna {

    private final int OVEN_TIME = 40;
    private final int MINUTES_PER_LAYER = 2;

    public int expectedMinutesInOven() {
        return OVEN_TIME;
    }

    public int remainingMinutesInOven(int minutesInOven) {
        return expectedMinutesInOven() - minutesInOven;
    }

    public int preparationTimeInMinutes(int layers) {
        return layers * MINUTES_PER_LAYER;
    }

    public int totalTimeInMinutes(int layers, int minutesInOven) {
        return preparationTimeInMinutes(layers) + minutesInOven;
    }
}
