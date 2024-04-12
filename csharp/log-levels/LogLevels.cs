using System;

static class LogLine
{
    public static string Message(string logLine)
    {
        // Start one character after the position of : in the logLine
        var startIndex = logLine.IndexOf(':') + 1;
        return logLine.Substring(startIndex).Trim();
    }

    public static string LogLevel(string logLine)
    {
        // Start one character after the position of [ in the logLine
        var startIndex = logLine.IndexOf('[') + 1;
        // End one character before the position of ] in the logLine
        var endIndex = logLine.IndexOf(']') - 1;
        return logLine.Substring(startIndex, endIndex).ToLower();
    }

    public static string Reformat(string logLine) => $"{Message(logLine)} ({LogLevel(logLine)})";
}
