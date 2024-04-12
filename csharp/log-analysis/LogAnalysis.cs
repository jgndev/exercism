using System;

public static class LogAnalysis 
{
    // Extension method that takes in some delimiter and returns the substring after the delimiter.
    public static string SubstringAfter(this string str, string delimiter)
    {
        if (string.IsNullOrEmpty(str))
        {
            return string.Empty;
        }
        
        var startIndex = str.IndexOf(delimiter, StringComparison.Ordinal);

        if (startIndex == -1)
        {
            return string.Empty;
        }
        
        return str.Substring(startIndex + delimiter.Length);
    }

    // Extension method that takes in two delimiters and returns the substring between them.
    public static string SubstringBetween(this string str, string startDelimiter, string endDelimiter)
    {
        var startIndex = str.IndexOf(startDelimiter, StringComparison.Ordinal) + startDelimiter.Length;
        var endIndex = str.IndexOf(endDelimiter, startIndex, StringComparison.Ordinal);
        
        return str.Substring(startIndex, endIndex - startIndex);
    }

    // Extension method intended to return the string in a log line without the log level.
    public static string Message(this string str) => str.SubstringAfter(":").Trim();

    // Extension method intended to return the string for the log level and discarding the rest of the line.
    public static string LogLevel(this string str) => str.SubstringBetween("[", "]");
}