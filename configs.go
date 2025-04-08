package logger

import (
	"github.com/nlgolib/logger/internal/file"
)

/*
	This configs is package-level configs.
*/

var isDebug = false

// SetDebug sets the debug mode.
// If debug is true, the logger will print debug messages.
func SetDebug(debug bool) {
	isDebug = debug
}

// SetLogFolder sets the log folder.
// If the folder is not set, the default folder is "./logs".
func SetLogFolder(folder string) {
	file.LogFolderPath = folder
}
