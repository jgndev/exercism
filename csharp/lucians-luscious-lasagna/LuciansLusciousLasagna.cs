class Lasagna
{
    private const int OvenTime = 40;
    private const int MinutesPerLayer = 2;
    
    public int ExpectedMinutesInOven() => OvenTime;

    public int RemainingMinutesInOven(int cookTime) 
        => OvenTime - cookTime;

    public int PreparationTimeInMinutes(int layers)
        => layers * MinutesPerLayer;

    public int ElapsedTimeInMinutes(int layers, int cookTime)
        => PreparationTimeInMinutes(layers) + cookTime;
}
