public class Lasagna {

    private final int OVEN_TIME = 40;

    // expectedMinutesInOven returns the number of minutes lasagna is expected to take.
    public int expectedMinutesInOven() {
        return OVEN_TIME;
    }

    // remainingMinutesInOven returns the minutes left to cook in the oven.
    public int remainingMinutesInOven(int cookTime) {
        return expectedMinutesInOven() - cookTime;
    }

    // preparationTimeInMinutes returns the minutes added for the layers.
    // Each layer is an additional 2 minutes.
    public int preparationTimeInMinutes(int layers) {
        return layers * 2;
    }

    // totalTimeInMinutes returns the minutes you've been working on the lasagna.
    public int totalTimeInMinutes(int layers, int cookTime) {
        return preparationTimeInMinutes(layers) + cookTime;
    }
}
