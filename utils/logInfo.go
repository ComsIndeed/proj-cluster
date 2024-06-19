package utils

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

// LogEvent function to log events to a file and console (adapted for Termux)
func LogEvent(runtimeID string, message string, startTime time.Time) error {
    // Termux-specific log directory
    logDir := filepath.Join(os.Getenv("HOME"), "AAA_WORKSPACE/logs/proj-cluster/") // Replace with your desired path if different

    // Ensure logs directory exists
    if err := os.MkdirAll(logDir, 0755); err != nil { // Create directory if not present
        return fmt.Errorf("error creating logs directory: %w", err)
    }

    // Construct log file path
    logFilePath := filepath.Join(logDir, runtimeID+".log")

    // Open log file (create if not exists, append if exists)
    logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("error opening log file: %w", err)
    }
    defer logFile.Close()

    // Get current time and calculate runtime duration
    now := time.Now()
    duration := time.Since(startTime) // startTime is defined elsewhere, globally

    // Format log message
    logMessage := fmt.Sprintf("%s | Runtime: %s | %s\n", now.Format("2006-01-02 15:04:05"), duration.Round(time.Second), message)

    // Log to file
    if _, err := logFile.WriteString(logMessage); err != nil {
        return fmt.Errorf("error writing to log file: %w", err)
    }

    // Log to console (Termux)
    fmt.Println(logMessage) // Use fmt.Println for Termux console

    return nil
}

