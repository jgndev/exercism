public class Lasagna {
    private final int ovenTime = 40;

    public int expectedMinutesInOven() {
        return ovenTime;
    }

    public int remainingMinutesInOven(int cookTime) {
        return expectedMinutesInOven() - cookTime;
    }

    public int preparationTimeInMinutes(int layers) {
        return layers * 2;
    }

    public int totalTimeInMinutes(int layers, int cookTime) {
        return preparationTimeInMinutes(layers) + cookTime;
    }
}
