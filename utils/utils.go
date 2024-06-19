package utils

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

// ConnectionInfo holds information about the current network connection.
type ConnectionInfo struct {
    SSID           string
    Carrier        string
    SignalStrength int
    ConnectionType string
}

// GetConnectionInfo returns information about the current network connection.
func GetConnectionInfo() (*ConnectionInfo, error) {
    info := &ConnectionInfo{}

    // Determine the connection type
    connectionType, err := getConnectionType()
    if err != nil {
        return nil, err
    }
    info.ConnectionType = connectionType

    // Depending on the connection type, gather additional info
    switch connectionType {
    case "wifi":
        ssid, err := getWifiSSID()
        if err != nil {
            return nil, err
        }
        signalStrength, err := getWifiSignalStrength()
        if err != nil {
            return nil, err
        }
        info.SSID = ssid
        info.SignalStrength = signalStrength

    case "mobile data":
        carrier, err := getMobileCarrier()
        if err != nil {
            return nil, err
        }
        signalStrength, err := getMobileSignalStrength()
        if err != nil {
            return nil, err
        }
        info.Carrier = carrier
        info.SignalStrength = signalStrength

    case "ethernet":
        info.SignalStrength = 100 // Ethernet typically has a stable connection
    }

    return info, nil
}

// getConnectionType determines the current connection type (wifi, mobile data, or ethernet).
func getConnectionType() (string, error) {
    out, err := exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE,STATE", "device").Output()
    if err != nil {
        return "", err
    }
    devices := strings.Split(strings.TrimSpace(string(out)), "\n")
    for _, device := range devices {
        fields := strings.Split(device, ":")
        if len(fields) >= 3 && fields[2] == "connected" {
            if fields[1] == "wifi" {
                return "wifi", nil
            }
            if fields[1] == "gsm" {
                return "mobile data", nil
            }
            if fields[1] == "ethernet" {
                return "ethernet", nil
            }
        }
    }
    return "", errors.New("no active network connection found")
}

// getWifiSSID returns the SSID of the connected Wi-Fi network.
func getWifiSSID() (string, error) {
    out, err := exec.Command("nmcli", "-t", "-f", "DEVICE,SSID,ACTIVE", "dev", "wifi").Output()
    if err != nil {
        return "", err
    }
    networks := strings.Split(strings.TrimSpace(string(out)), "\n")
    for _, network := range networks {
        fields := strings.Split(network, ":")
        if len(fields) >= 3 && fields[2] == "yes" {
            return fields[1], nil
        }
    }
    return "", errors.New("not connected to any Wi-Fi network")
}

// getWifiSignalStrength returns the signal strength of the connected Wi-Fi network.
func getWifiSignalStrength() (int, error) {
    out, err := exec.Command("nmcli", "-t", "-f", "DEVICE,SIGNAL,ACTIVE", "dev", "wifi").Output()
    if err != nil {
        return 0, err
    }
    networks := strings.Split(strings.TrimSpace(string(out)), "\n")
    for _, network := range networks {
        fields := strings.Split(network, ":")
        if len(fields) >= 3 && fields[2] == "yes" {
            signalStrength, err := strconv.Atoi(fields[1])
            if err != nil {
                return 0, err
            }
            return signalStrength, nil
        }
    }
    return 0, errors.New("not connected to any Wi-Fi network")
}

// getMobileCarrier returns the carrier name of the connected mobile data network.
func getMobileCarrier() (string, error) {
    out, err := exec.Command("nmcli", "-t", "-f", "GENERAL.DEVICE,GENERAL.CONNECTION,GENERAL.TYPE,GENERAL.STATE", "device").Output()
    if err != nil {
        return "", err
    }
    devices := strings.Split(strings.TrimSpace(string(out)), "\n")
    for _, device := range devices {
        fields := strings.Split(device, ":")
        if len(fields) >= 4 && fields[3] == "connected" && fields[2] == "gsm" {
            return fields[1], nil
        }
    }
    return "", errors.New("not connected to any mobile data network")
}

// getMobileSignalStrength returns the signal strength of the connected mobile data network.
func getMobileSignalStrength() (int, error) {
    out, err := exec.Command("nmcli", "-t", "-f", "DEVICE,SIGNAL", "device", "gsm").Output()
    if err != nil {
        return 0, err
    }
    devices := strings.Split(strings.TrimSpace(string(out)), "\n")
    for _, device := range devices {
        fields := strings.Split(device, ":")
        if len(fields) >= 2 {
            signalStrength, err := strconv.Atoi(fields[1])
            if err != nil {
                return 0, err
            }
            return signalStrength, nil
        }
    }
    return 0, errors.New("not connected to any mobile data network")
}
