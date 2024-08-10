public class LogLevels {
    
    public static String message(String logLine) {
        // the second element in the array has the message
        // return it after trimming whitespace
        return splitLine(logLine)[1].trim();
    }

    public static String logLevel(String logLine) {
        // gets a string like "[WARNING]:"
        var level = splitLine(logLine)[0].trim();
        // return a substring like "warning"
        return level.substring(1, level.length() - 1).toLowerCase();
    }

    public static String reformat(String logLine) {
        return String.format("%s (%s)", message(logLine), logLevel(logLine));
    }

    private static String[] splitLine(String line) {
        if (!line.contains(":")) {
            throw new IllegalArgumentException("Invalid log format: missing colon in string");
        }
        return line.split(":" );
    }
}
