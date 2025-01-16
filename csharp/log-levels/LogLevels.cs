using System;

static class LogLine
{
    public static string Message(string logLine)
        => logLine.Split(':')[1].Trim();


    public static string LogLevel(string logLine)
    {
        // char position after the leading left bracket '['
        var startIndex = logLine.IndexOf('[') + 1;
        // char position before the closing right bracket ']'
        var endIndex = logLine.IndexOf(']') - 1;
        // take a substring from startIndex to endIndex and lowercase it
        return logLine.Substring(startIndex, endIndex).ToLower();
    }

    public static string Reformat(string logLine)
        => $"{Message(logLine)} ({LogLevel(logLine)})";
}
