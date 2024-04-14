static class AssemblyLine
{

    private const int CarsPerHour = 221;
    private const int MinutesPerHour = 60;
    private const double NoErrorRate = 1.0;
    private const double LowErrorRate = 0.9;
    private const double MediumErrorRate = 0.8;
    private const double HighErrorRate = 0.77;
    
    public static double SuccessRate(int speed)
    {
        // Treat negative speeds as 0.
        if (speed < 0)
        {
            return 0;
        }
        
        return speed switch
        {
            0 => 0,
            <= 4 and >= 1 => NoErrorRate,
            <= 8 and >= 5 => LowErrorRate,
            9 => MediumErrorRate,
            _ => HighErrorRate,
        };
    }
    
    public static double ProductionRatePerHour(int speed)
        => CarsPerHour * speed * SuccessRate(speed);

    public static int WorkingItemsPerMinute(int speed)
    {
        var perHour = ProductionRatePerHour(speed);
        return (int)(perHour / MinutesPerHour);
    }
}
