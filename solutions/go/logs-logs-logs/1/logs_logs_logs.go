package logs

import "strings"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
    app := "default"
    match := 0
	for _, char := range log {
        switch char {
            case '❗':
            	app = "recommendation"
            	match++
            	break
            case '🔍':
                app = "search"
            	match++
            	break
            case '☀':
                app = "weather"
            	match++
                break
        }
        if match > 0 {
            break
        }
    }
    return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
