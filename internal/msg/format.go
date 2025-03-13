package msg

import (
	"fmt"
	"time"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"

	// ANSI color codes
	colorReset  = "\033[0m"
	colorGray   = "\033[90m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorBlue   = "\033[34m"
)

func FormatMessage(level string, messages ...any) string {
	color := getColorForLevel(level)
	return fmt.Sprintf("%s[%s]%s %s[%s]%s %v",
		color, level, colorReset,
		colorGray, getNowTime(), colorReset,
		fmt.Sprint(messages...))
}

func getColorForLevel(level string) string {
	switch level {
	case LevelDebug:
		return colorBlue
	case LevelInfo:
		return colorGreen
	case LevelWarn:
		return colorYellow
	case LevelError:
		return colorRed
	default:
		return colorReset
	}
}

func getNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
