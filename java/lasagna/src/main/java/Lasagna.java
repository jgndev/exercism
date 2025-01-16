public class Lasagna {
    private final int OvenTime = 40;
    private final int MinutesPerLayer = 2;

    public int expectedMinutesInOven() {
        return OvenTime;
    }

    public int remainingMinutesInOven(int minutesCooked) {
        return expectedMinutesInOven() - minutesCooked;
    }

    public int preparationTimeInMinutes(int layers) {
       return layers * MinutesPerLayer;
    }

    public int totalTimeInMinutes(int layers, int minutesCooked) {
        return preparationTimeInMinutes(layers) + minutesCooked;
    }
}
