package logger

import (
	"github.com/nlgolib/logger/internal/file"
)

var (
	DefaultLogFolder = file.DefaultLogFolder
	EnvLogFolderPath = file.EnvLogFolderPath
	IsDebug          = false
)

func SetLogFolder(folder string) {
	file.DefaultLogFolder = folder
	DefaultLogFolder = folder
}
