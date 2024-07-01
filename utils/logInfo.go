package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// LogEvent logs a formatted message to both the console and a specified log file.
func LogEvent(runtimeID string, startTime time.Time) error {
    // 1. Log File Setup (Improved Error Handling)
    logDir := filepath.Join(os.Getenv("HOME"), "AAA_WORKSPACE/logs/proj-cluster/")
    if err := os.MkdirAll(logDir, 0755); err != nil {
        return fmt.Errorf("error creating logs directory: %w", err)
    }

    logFilePath := filepath.Join(logDir, runtimeID+".log")
    logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("error opening log file: %w", err)
    }
    defer logFile.Close() // Ensure file is closed even if there's an error later

    // LOG MESSAGES
    
    var logMessage strings.Builder
    
    now := time.Now()
    duration := time.Since(startTime).Round(time.Second)
    logMessage.WriteString(fmt.Sprintf("\033[36m<>\033[32m %s | Runtime: %s\033[0m\n",
    now.Format("2006-01-02 15:04:05"), duration,))

    logMessage.WriteString("Devices:\n")
    logMessage.WriteString("- \"FakePhone 2410S\" | LastBeat: 5s | Services: [Bot, Bot1, Bot3]\n")
    logMessage.WriteString("- \"Cailiber 5A\" | LastBeat: 3s | Services: [Bot3]\n")
    logMessage.WriteString("- \"[IDLE] aPhone 250S\" | LastBeat: 6s | Services: []\n")

    fmt.Println(logMessage.String())
    if _, err := logFile.WriteString(logMessage.String()); err != nil {
        // Log to console if there was an error writing to the file
        fmt.Fprintf(os.Stderr, "Error writing to log file: %v\n", err) // Print to stderr for error logs
    }

    return nil 
}
