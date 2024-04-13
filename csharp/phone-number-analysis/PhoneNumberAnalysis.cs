using System;

public static class PhoneNumber
{

    private const string NyAreaCode = "212";
    private const string FakeExchangeCode = "555";
    private const int PhoneNumberLength = 12;
    
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber)
    {
        if (phoneNumber.Length != PhoneNumberLength)
        {
            throw new ArgumentException("Phone number must be in the format NNN-NNN-NNNN");
        }

        var areaCode = phoneNumber.Substring(0, 3);
        var exchangeCode = phoneNumber.Substring(4, 3);
        var localNumber = phoneNumber.Substring(8, 4);

        var isNewYork = areaCode == NyAreaCode;
        var isFake = exchangeCode == FakeExchangeCode;

        return (IsNewYork: isNewYork, IsFake: isFake, LocalNumber: localNumber);
    }

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo) 
        => phoneNumberInfo.IsFake;
}
