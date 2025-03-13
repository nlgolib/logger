package logger

import (
	"github.com/nlgolib/logger/internal/file"
)

var (
	DefaultLogFolder = file.DefaultLogFolder
	EnvLogFolderPath = file.EnvLogFolderPath
	IsDebug          = false
)
