class Lasagna
{
    private const int OvenTime = 40;
    
    public int ExpectedMinutesInOven() 
        => OvenTime;

    public int RemainingMinutesInOven(int cookTime) 
        => OvenTime - cookTime;

    public int PreparationTimeInMinutes(int layers) 
        => layers * 2;

    public int ElapsedTimeInMinutes(int layers, int cookTime) 
        => PreparationTimeInMinutes(layers) + cookTime;
}
