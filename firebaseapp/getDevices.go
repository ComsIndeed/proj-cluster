package firebaseapp

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"firebase.google.com/go/db"
)

// GetDevices retrieves device information from the Firebase Realtime Database.
func GetDevices(ctx context.Context, client *db.Client) (map[string]interface{}, error) {
    ref := client.NewRef("deviceStatuses")

    var devices map[string]interface{}
    if err := ref.Get(ctx, &devices); err != nil {
        return nil, fmt.Errorf("error retrieving devices: %v", err)
    }
    return devices, nil
}

// FormatAndPrintDevices formats and prints the device information, including elapsed time.
func FormatAndPrintDevices(devices map[string]interface{}) {
    // Sorting device IDs for consistent output
    ids := make([]string, 0, len(devices))
    for id := range devices {
        ids = append(ids, id)
    }
    sort.Strings(ids)

    location, err := time.LoadLocation("Asia/Manila")
    if err != nil {
        fmt.Println("Error loading location:", err)
        return
    }

    fmt.Println("DEVICES:")
    for _, id := range ids {
        deviceInfo := devices[id].(map[string]interface{})

        name, _ := deviceInfo["deviceId"].(string)
        os, _ := deviceInfo["operatingSystem"].(string)
        servicesInterface, _ := deviceInfo["services"].([]interface{})
        lastBeatTime := time.Unix(int64(deviceInfo["time"].(float64)), 0).In(location) // Time in your location
        elapsedTime := time.Since(lastBeatTime).Round(time.Second)

        // Convert services from []interface{} to []string
        var services []string
        for _, svc := range servicesInterface {
            services = append(services, fmt.Sprint(svc))
        }
        servicesStr := strings.Join(services, ", ")

        // Print formatted device info
        fmt.Printf("- \"%s\" | OS: %s | Last Beat: %s ago | Services: [%s]\n", name, os, elapsedTime, servicesStr)
    }
}

