using System;
using System.Runtime.CompilerServices;

public static class LogAnalysis 
{
    public static string SubstringAfter(this string str, string delimiter)
    {
        if (string.IsNullOrEmpty(str))
        {
            return string.Empty;
        }

        var index = str.IndexOf(delimiter, StringComparison.Ordinal);

        return str.Substring(index + delimiter.Length);
    }

    public static string SubstringBetween(this string str, string firstDelimiter, string secondDelimiter)
    {
        var startIndex = str.IndexOf(firstDelimiter, StringComparison.Ordinal) + firstDelimiter.Length;
        var endIndex = str.IndexOf(secondDelimiter, startIndex, StringComparison.Ordinal);

        if (startIndex < 0 || endIndex < 0 || endIndex <= startIndex)
        {
            return string.Empty;
        }

        return str.Substring(startIndex, endIndex - startIndex);
    }
    
    public static string Message(this string str) 
        => str.SubstringAfter(":").Trim();

    public static string LogLevel(this string str) 
        => str.SubstringBetween("[", "]");
}