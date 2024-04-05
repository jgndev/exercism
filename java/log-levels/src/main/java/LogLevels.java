public class LogLevels {

    // Return the message in the log line with the level removed
    // "[ERROR]: Invalid operation" returns "Invalid Operation"
    public static String message(String logLine) {
        var colonIndex = getColonIndex(logLine);

        return logLine.substring(colonIndex + 1).trim();
    }

    // Return the log level in lowercase.
    // "[ERROR]: Invalid Operation" returns "error"
    public static String logLevel(String logLine) {
        var colonIndex = getColonIndex(logLine);
        return logLine.substring(1, colonIndex - 1).trim().toLowerCase();
    }

    // Return the message with the log level at the end of the message inside parenthesis lowercase.
    // "[ERROR]: Invalid Operation" returns "Invalid operation (error)"
    public static String reformat(String logLine) {
        var level = logLevel(logLine);
        var message = message(logLine);

        return message + " (" + level + ")";
    }

    // Return the index of the ':' character in the string for use with substr String method.
    private static int getColonIndex(String line) {
        return line.indexOf(":");
    }
}
