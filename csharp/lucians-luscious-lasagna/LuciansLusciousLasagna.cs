class Lasagna {
    private const int OVEN_TIME = 40; 
        
    public int ExpectedMinutesInOven() {
        return OVEN_TIME;
    } 

    public int RemainingMinutesInOven(int cookTime) {
        return OVEN_TIME - cookTime;
    }

    public int PreparationTimeInMinutes(int layers) {
        return layers * 2;
    }

    public int ElapsedTimeInMinutes(int layers, int cookTime) {
        return PreparationTimeInMinutes(layers) + cookTime;
    }
}
