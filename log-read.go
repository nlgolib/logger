package logger

import (
	"github.com/nlgolib/logger/internal/file"
	"github.com/nlgolib/logger/internal/msg"
)

const (
	LevelDebug = msg.LevelDebug
	LevelInfo  = msg.LevelInfo
	LevelWarn  = msg.LevelWarn
	LevelError = msg.LevelError
)

func ListLogFiles() []string {
	return file.ListLogFiles()
}

func ReadLogFile(date string, levels ...string) string {
	return file.ReadLogFile(date, levels...)
}

func DeleteLogFile(name string) {
	file.DeleteLog(name)
}
