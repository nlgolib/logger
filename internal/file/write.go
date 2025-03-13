package file

import (
	"fmt"
	"os"
	"time"

	"github.com/nlgolib/env"
)

const DefaultLogFolder = "./logs"

var EnvLogFolderPath = env.New("LOG_FOLDER_PATH")

func WriteLog(message string) {
	filePath := getLogFilePath()
	os.MkdirAll(getLogFolderPath(), os.ModePerm)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("Error writing log file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(message + "\n")
	if err != nil {
		fmt.Println("Error writing log file:", err)
	}
}

func DeleteLog(name string) {
	filePath := getLogFolderPath() + "/" + name
	os.Remove(filePath)
}

func getTodayLogFileName() string {
	return time.Now().Format("2006-01-02") + ".log"
}

func getLogFolderPath() string {
	appFolder := EnvLogFolderPath.String()
	if appFolder == "" {
		appFolder = DefaultLogFolder
	}
	return appFolder
}

func getLogFilePath() string {
	return getLogFolderPath() + "/" + getTodayLogFileName()
}
