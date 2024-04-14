using System;

static class SavingsAccount
{
    public static float InterestRate(decimal balance)
    {
        var interest = 0.0f;

        switch (balance)
        {
           case < 0:
               interest = 3.213f;
               break;
           
           case < 1000 and >= 0:
               interest = 0.5f;
               break;
           
           case < 5000 and >= 1000:
               interest = 1.621f;
               break;
           
           default:
               interest = 2.475f;
               break;
        }
        

        return interest;
    }

    public static decimal Interest(decimal balance)
    {
        var interestRate = InterestRate(balance);
        return balance * (decimal)interestRate / 100;
    }

    public static decimal AnnualBalanceUpdate(decimal balance)
        => balance + Interest(balance);

    public static int YearsBeforeDesiredBalance(decimal balance, decimal targetBalance)
    {
        var currentBalance = balance;
        var years = 0;

        while (currentBalance < targetBalance)
        {
            currentBalance += Interest(currentBalance);
            years++;
        }

        return years;
    }
}
