using System;

static class LogLine
{
    public static string Message(string logLine)
    {
        var index = logLine.IndexOf(':');
        // The string we want begins after the colon and a space.
        return logLine.Substring(index + 2).Trim();
    }

    public static string LogLevel(string logLine)
    {
        var startIndex = logLine.IndexOf('[');
        var endIndex = logLine.IndexOf(']');

        // Beginning from the first character after the '[' and ending with the distance between the start end ending index
        return logLine.Substring(startIndex + 1, endIndex - startIndex - 1).ToLower().Trim();
    }

    public static string Reformat(string logLine) 
        => $"{Message(logLine)} ({LogLevel(logLine)})";
}
