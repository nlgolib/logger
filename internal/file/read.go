package file

import (
	"os"
	"strings"
)

func ListLogFiles() []string {
	files, err := os.ReadDir(getLogFolderPath())
	if err != nil {
		return nil
	}
	names := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			names = append(names, file.Name())
		}
	}
	return names
}

func ReadLogFile(name string, levels ...string) string {
	filePath := getLogFolderPath() + "/" + name
	content, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}

	if len(levels) == 0 {
		return string(content)
	}

	lines := strings.Split(string(content), "\n")
	output := make([]string, 0)
	open := false
	for _, line := range lines {
		if toCheck(line) {
			open = isLevel(line, levels)
		}
		if open {
			output = append(output, line)
		}
	}
	return strings.Join(output, "\n")
}

func toCheck(line string) bool {
	return strings.Contains(line, "[DEBUG]") ||
		strings.Contains(line, "[INFO]") ||
		strings.Contains(line, "[WARN]") ||
		strings.Contains(line, "[ERROR]")
}

func isLevel(line string, levels []string) bool {
	for _, level := range levels {
		if strings.Contains(line, "["+level+"]") {
			return true
		}
	}
	return false
}
