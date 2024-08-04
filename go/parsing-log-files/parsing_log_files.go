package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	validPrefixes := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`"[^"]*password[^"]*"`)
	for _, line := range lines {
		if re.MatchString(strings.ToLower(line)) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))
	re := regexp.MustCompile(`User\s+(\S+)`)

	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			result[i] = "[USR] " + match[1] + " " + line
		} else {
			result[i] = line
		}
	}

	return result
}
