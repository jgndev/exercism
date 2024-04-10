using System;

public static class PhoneNumber
{

    private const string NewYorkAreaCode = "212";
    private const string FakeExchangeCode = "555";
    
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber)
    {

        if (string.IsNullOrWhiteSpace(phoneNumber) || phoneNumber.Length != 12)
        {
            throw new ArgumentException("Phone number must be in the format 'XXX-XXX-XXXX'", nameof(phoneNumber));
        }
        
        var areaCode = phoneNumber.Substring(0, 3);
        var newYorkAreaCode = areaCode == NewYorkAreaCode;

        var exchangeCode = phoneNumber.Substring(4, 3);
        var fakeExchangeCode = exchangeCode == FakeExchangeCode;

        var lineNumber = phoneNumber.Substring(8, 4);

        return (IsNewYork: newYorkAreaCode, IsFake: fakeExchangeCode, LocalNumber: lineNumber);
    }

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo) 
        => phoneNumberInfo.IsFake;
}
