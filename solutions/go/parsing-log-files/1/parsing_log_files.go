package parsinglogfiles

import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	validLog := regexp.MustCompile(`^[[TRC|DBG|INF|WRN|ERR|FTL]+].*`)
    return validLog.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
    count := 0
    for _, v := range lines {
        if re.FindString(v) != "" {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    keepSlice := re.Split(text, -1)
    var fixedString string
    for _, v := range keepSlice {
        fixedString = fixedString + v
    }
    return fixedString
}

func TagWithUserName(lines []string) []string {
    var newSlice []string
	re := regexp.MustCompile(`User\s+(\w+)`)
    for _, v := range lines {
        match := re.FindStringSubmatch(v)
        if match != nil {
            newSlice = append(newSlice, fmt.Sprintf("[USR] %s %s", match[1], v))
        } else {
            newSlice = append(newSlice, v)
        }
    }
    return newSlice
}
