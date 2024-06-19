package utils

import (
    "os/exec"
    "regexp"
    "strings"
)

// ConnectionInfo struct to hold connection details
type ConnectionInfo struct {
    SSID           string `json:"ssid,omitempty"` // Optional SSID
    Carrier        string `json:"carrier,omitempty"` // Optional carrier
    SignalStrength string `json:"signalStrength,omitempty"` // Optional signal strength
    ConnectionType string `json:"connectionType,omitempty"` // Optional connection type
}

// GetConnectionInfo retrieves connection information.
func GetConnectionInfo() (*ConnectionInfo, error) {
    cmd := exec.Command("iwconfig")
    output, err := cmd.Output()

    if err != nil {
        // Optional: Log error for debugging
        //fmt.Println("Error running iwconfig:", err)
        return nil, err
    }

    lines := strings.Split(string(output), "\n")
    info := &ConnectionInfo{} // Initialize ConnectionInfo struct
    for _, line := range lines {
        if strings.Contains(line, "ESSID:") {
            info.ConnectionType = "Wi-Fi"
            re := regexp.MustCompile(`ESSID:"([^"]*)"`)
            match := re.FindStringSubmatch(line)
            if len(match) > 1 {
                info.SSID = match[1]
            }
            break
        }
    }
    
    if info.ConnectionType == "" {
        info.ConnectionType = "unknown"
    }

    if info.SSID == "" {
        info.SSID = "No WiFi Connection"
    }

    return info, nil
}

