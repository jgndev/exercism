class Lasagna
{
    private static int OvenTime = 40;
    
    public int ExpectedMinutesInOven()
    {
        return OvenTime;
    }

    public int RemainingMinutesInOven(int ovenMinutes)
    {
        return OvenTime - ovenMinutes;
    } 

    public int PreparationTimeInMinutes(int layers)
    {
        return layers * 2;
    }

    public int ElapsedTimeInMinutes(int layers, int ovenTime)
    {
        return ovenTime + PreparationTimeInMinutes(layers);
    }
}
