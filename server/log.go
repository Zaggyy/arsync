package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

const LOG_FOLDER_NAME = "logs"

func PrepareLogs() {
	// Create the logs folder if it doesn't exist
	if _, err := os.Stat(LOG_FOLDER_NAME); os.IsNotExist(err) {
		os.Mkdir(LOG_FOLDER_NAME, 0755)
	}

	var latestLogPath string = path.Join(LOG_FOLDER_NAME, "latest.log") // logs/latest.log

	// Check if latest.log exists
	if _, err := os.Stat(latestLogPath); err == nil {
		// Rename it to a timestamped log
		fStat, _ := os.Stat(latestLogPath)
		lastModified := fStat.ModTime().Format("2006-01-02_15-04-05")
		lastModifiedLogPath := path.Join(LOG_FOLDER_NAME, lastModified+".log") // logs/2021-01-01_12-00-00.log

		os.Rename(latestLogPath, path.Join(LOG_FOLDER_NAME, lastModifiedLogPath))

		// Compress the log
		originalLog, _ := os.Open(lastModifiedLogPath)
		defer originalLog.Close()

		compressedLog, _ := os.Create(lastModifiedLogPath + ".gz")
		defer compressedLog.Close()

		// Create a gzip writer
		w := gzip.NewWriter(compressedLog)
		defer w.Close()

		io.Copy(w, originalLog)
		w.Flush()

		// Remove the original log
		os.Remove(latestLogPath)
	}

	// Create the new latest.log
	os.Create(latestLogPath)
}

func Log(message string, level string) {
	log.Printf(fmt.Sprintf("%s: %s", level, message))
	f, _ := os.OpenFile("logs/latest.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	log.SetOutput(f)
	log.Printf(fmt.Sprintf("%s: %s", level, message))
	log.SetOutput(os.Stdout)
}
