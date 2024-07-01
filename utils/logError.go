package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogError logs an error message with timestamp and runtime ID to both the console and log file.
func LogError(runtimeID string, err error) {
    // 1. Log File Setup (Error Handling)
    logDir := filepath.Join(os.Getenv("HOME"), "AAA_WORKSPACE/logs/proj-cluster/")
    if err := os.MkdirAll(logDir, 0755); err != nil {
        // If creating the directory fails, log the error to stderr and return
        fmt.Fprintf(os.Stderr, "Error creating logs directory: %v\n", err)
        return 
    }

    logFilePath := filepath.Join(logDir, runtimeID+".log")
    logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        // If opening the log file fails, log the error to stderr and return
        fmt.Fprintf(os.Stderr, "Error opening log file: %v\n", err)
        return
    }
    defer logFile.Close()

    // 2. Log Message Formatting (Error Message, Timestamp, and Runtime ID)
		location, locationError := time.LoadLocation("Asia/Manila")
		if locationError != nil {
			fmt.Printf("Error getting location: %v", locationError)
		}
    timestamp := time.Now().In(location).Format("2006-01-02 15:04:05") // Time in your location
    logMessage := fmt.Sprintf("\033[31m[Error] %s | Runtime ID: %s | %v\033[0m\n", timestamp, runtimeID, err) // Red color for errors

    // 3. Logging (Combined Output)
    fmt.Fprintln(os.Stderr, logMessage) // Print to stderr for error logs

    if _, err := logFile.WriteString(logMessage); err != nil {
        // If writing to file fails, log the error to stderr 
        fmt.Fprintf(os.Stderr, "Error writing to log file: %v\n", err)
    }
}
