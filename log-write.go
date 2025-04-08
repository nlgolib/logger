package logger

import (
	"fmt"

	"github.com/nlgolib/logger/internal/file"
	"github.com/nlgolib/logger/internal/msg"
)

// Debug logs a debug message.
// If debug is true, the message will be printed to the console.
// The logs will NOT be written to the file.
func Debug(messages ...any) {
	if isDebug {
		fmt.Println(msg.FormatMessage(msg.LevelDebug, messages...))
	}
}

// Info logs an info message.
// The message will be printed to the console and written to the file.
func Info(messages ...any) {
	msg := msg.FormatMessage(msg.LevelInfo, messages...)
	fmt.Println(msg)
	file.WriteLog(msg)
}

// Warn logs a warning message.
// The message will be printed to the console and written to the file.
func Warn(messages ...any) {
	msg := msg.FormatMessage(msg.LevelWarn, messages...)
	fmt.Println(msg)
	file.WriteLog(msg)
}

// Error logs an error message.
// The message will be printed to the console and written to the file.
func Error(messages ...any) {
	msg := msg.FormatMessage(msg.LevelError, messages...)
	fmt.Println(msg)
	file.WriteLog(msg)
}
