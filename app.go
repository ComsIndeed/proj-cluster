package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	// "strings"
	"time"

	"github.com/google/uuid"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type ConnectionInfo struct {
	Type            string `json:"type"`
	Connectivity    int    `json:"connectivity"`
	SSID            string `json:"ssid"`
	SignalLeveldBm int    `json:"signal_level_dBm"`
}

// getNetworkInterfaceName gets the main network interface name.
func getNetworkInterfaceName() (string, error) {
    // TODO: Implement the logic to get the main network interface name
    // This will vary depending on your operating system and network setup
    // You might use commands like "ip", "route", or platform-specific libraries
	return "wlan0", nil // Placeholder for now
}

// getWifiSignalStrength retrieves Wi-Fi signal strength using iwconfig
func getWifiSignalStrength(iface string) (ConnectionInfo, error) {
	cmd := exec.Command("iwconfig", iface)
	output, err := cmd.Output()
	if err != nil {
		return ConnectionInfo{}, fmt.Errorf("error running iwconfig: %v", err)
	}

	// Parsing the output from iwconfig to get Wi-Fi details
	reSignal := regexp.MustCompile(`Signal level=(-?\d+) dBm`)
	matchSignal := reSignal.FindStringSubmatch(string(output))
	signalLevel := 0
	if len(matchSignal) > 1 {
		signalLevel, _ = strconv.Atoi(matchSignal[1])
	}

	reSSID := regexp.MustCompile(`ESSID:"(.*?)"`)
	matchSSID := reSSID.FindStringSubmatch(string(output))
	ssid := ""
	if len(matchSSID) > 1 {
		ssid = matchSSID[1]
	}

	reLinkQuality := regexp.MustCompile(`Link Quality=(\d+)/\d+`)
	matchLinkQuality := reLinkQuality.FindStringSubmatch(string(output))
	strengthPercent := 0
	if len(matchLinkQuality) > 1 {
		linkQuality, _ := strconv.Atoi(matchLinkQuality[1])
		strengthPercent = linkQuality * 100 / 70 
	}

	return ConnectionInfo{
		Type:            "Wi-Fi", 
		Connectivity:    strengthPercent,
		SSID:            ssid,
		SignalLeveldBm: signalLevel, 
	}, nil
}



func sendHeartbeat(ref *db.Ref, connectionID string, status map[string][]string) {
	iface, err := getNetworkInterfaceName()
	if err != nil {
		fmt.Println("Error getting network interface name:", err)
		return
	}

	connInfo, err := getWifiSignalStrength(iface)
	if err != nil {
		fmt.Println("Error getting Wi-Fi signal strength:", err)
		return
	}

	hostname, _ := os.Hostname() 
	payload := map[string]interface{}{
		"device":    runtime.GOOS + " " + runtime.GOARCH,
		"hostname":  hostname,
		"heartbeat": time.Now().Unix(),
		"services":  status["services"],
		"connection": map[string]interface{}{
			"type":         connInfo.Type,
			"connectivity": connInfo.Connectivity,
			"ssid":         connInfo.SSID,
			"signal_level_dBm": connInfo.SignalLeveldBm,
		},
	}

	err = ref.Child("devices").Child(connectionID).Set(context.Background(), payload)
	if err != nil {
		log.Fatal("Error setting device data:", err)
	}
}

func main() {
	// Get executable directory and credentials file path
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Error getting executable path:", err)
	}
	exeDir := filepath.Dir(exePath)
	credentialsFile := filepath.Join(exeDir, "project-academic-weapon-firebase-adminsdk-irvv8-de78853ef1.json")

	// Check if credentials file exists
	if _, err := os.Stat(credentialsFile); err != nil {
		log.Fatal("Credentials file not found in the same directory as the executable:", err)
	}

	// Initialize Firebase App
	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Error initializing Firebase app:", err)
	}

	// Get a reference to the Realtime Database
	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatal("Error getting database client:", err)
	}
	dbRef := client.NewRef("devices")

	connectionID := uuid.New().String()
	status := map[string][]string{"services": {}} 

	for len(status["services"]) == 0 { 
		sendHeartbeat(dbRef, connectionID, status)
		time.Sleep(5 * time.Second)
	}

	for len(status["services"]) > 0 {
		sendHeartbeat(dbRef, connectionID, status)
		time.Sleep(1 * time.Second)
	}
}

