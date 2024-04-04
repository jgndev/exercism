class Lasagna
{

    private const int OVEN_TIME = 40;

    // Returns how many minutes the lasagna
    public int ExpectedMinutesInOven() => OVEN_TIME;

    // Returns the remaining minutes in the oven
    public int RemainingMinutesInOven(int cookTime) => OVEN_TIME - cookTime;

    // Returns the additional time needed per layer. Each layer requires 2 minutes.
    public int PreparationTimeInMinutes(int layers) => layers * 2;

    // Returns how many minutes you've been working on the lasagna, which is the
    // sum of the preparation time and minutes the lasagna has been in the oven.
    public int ElapsedTimeInMinutes(int layers, int cookTime)
    {
        return PreparationTimeInMinutes(layers) + cookTime;
    }
}
