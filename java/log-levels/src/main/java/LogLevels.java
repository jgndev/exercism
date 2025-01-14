public class LogLevels {

    // message returns the message of the logline with the loglevel removed.
    public static String message(String logLine) {
        return logLine.split(":", 2)[1].trim();
    }

    // logLevel returns the log level in lowercase.
    public static String logLevel(String logLine) {
        return logLine.substring(1, logLine.indexOf("]")).toLowerCase();
    }

    // reformat returns the logLine in a new format.
    public static String reformat(String logLine) {
        return message(logLine) + " (" + logLevel(logLine) + ")";
    }
}
