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

var Reader = &reader{}

type reader struct{}

// ListFiles lists the log files.
func (r *reader) ListFiles() []string {
	return file.ListLogFiles()
}

// Read reads the log file.
// You can specify the date and levels of the log file.
func (r *reader) Read(date string, levels ...string) string {
	return file.ReadLogFile(date, levels...)
}

// DeleteLogFile deletes the log file.
func (r *reader) Delete(name string) {
	file.DeleteLog(name)
}
