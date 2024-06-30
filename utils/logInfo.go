package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func LogEvent(runtimeID string, startTime time.Time) error {

	// For the directory
	logDir := filepath.Join(os.Getenv("HOME"), "AAA_WORKSPACE/logs/proj-cluster/")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("error creating logs directory: %w", err)
	}

	// For the file
	logFilePath := filepath.Join(logDir, runtimeID+".log")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening log file: %w", err)
	}
	defer logFile.Close()

	// THESE ARE THE ONES GATHERING DATA

	// Get current time and calculate runtime duration
	now := time.Now()
	duration := time.Since(startTime)

	// Format log message
	logMessage := fmt.Sprintf("\033[36m<>\033[32m %s | Runtime: %s\033[0m\nFirebase: (Connected | Not connected)\nServices: (List)\n", now.Format("2006-01-02 15:04:05"), duration.Round(time.Second))
	if _, err := logFile.WriteString(logMessage); err != nil {
		return fmt.Errorf("error writing to log file: %w", err)
	}
	fmt.Println(logMessage)

	return nil
}
