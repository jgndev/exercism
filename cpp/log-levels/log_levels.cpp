#include <string>

namespace log_line {
    std::string message(std::string line) {
        // Example log line: "[INFO]: some message"
        auto idx = line.find(": ");
        return std::string(line.substr(idx + 2));
    }

    std::string log_level(std::string line) {
        // Skip the '[' and get the substring until ']'
        auto level_end = line.find(']') - 1;
        return std::string(line.substr(1, level_end));
    }

    std::string reformat(std::string line) {
        std::string result;
        // pre-allocate the space for the new message
        result.reserve(line.length() + 2);

        auto msg = message(line);
        auto level = log_level(line);

        result.append(msg);
        result.append(" (");
        result.append(level);
        result.append(")");

        return result;
    }
}
