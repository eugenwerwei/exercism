package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(\W)*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password"`)
	count := 0

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line+\d*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User*\s+(\w+)`)

	for i, line := range lines {
		str := re.FindStringSubmatch(line)
		if str != nil {
			lines[i] = fmt.Sprintf("[USR] %s %s", str[1], line)
		}
	}

	return lines
}
