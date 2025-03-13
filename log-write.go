package logger

import (
	"fmt"

	"github.com/nlgolib/logger/internal/file"
	"github.com/nlgolib/logger/internal/msg"
)

func Debug(messages ...any) {
	if IsDebug {
		fmt.Println(msg.FormatMessage(msg.LevelDebug, messages...))
	}
}

func Info(messages ...any) {
	msg := msg.FormatMessage(msg.LevelInfo, messages...)
	fmt.Println(msg)
	file.WriteLog(msg)
}

func Warn(messages ...any) {
	msg := msg.FormatMessage(msg.LevelWarn, messages...)
	fmt.Println(msg)
	file.WriteLog(msg)
}

func Error(messages ...any) {
	msg := msg.FormatMessage(msg.LevelError, messages...)
	fmt.Println(msg)
	file.WriteLog(msg)
}
