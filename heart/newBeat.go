package heart

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"firebase.google.com/go/db"
)

func Heartbeat(runtimeId string, ctx context.Context, databaseClient *db.Client) {
	var devicesRef = databaseClient.NewRef("deviceStatuses")

	var hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	// var connectionInfo, err2 = utils.GetConnectionInfo()
	// if err2 != nil {
	// 	fmt.Printf("CONNECTION INFO ERROR: %v", err2)
	// }

	var dataOut = map[string]interface{}{
		"id":              runtimeId,
		"deviceId":        hostname,
		"operatingSystem": runtime.GOOS,
		// "connection": map[string]interface{}{
		// 	"ssid":           connectionInfo.SSID,
		// 	"carrier":        connectionInfo.Carrier,
		// 	"signalStrength": connectionInfo.SignalStrength,
		// 	"connectionType": connectionInfo.ConnectionType,
		// },
		"time":     time.Now().Unix(),
		"services": []string{""},
	}

	go func() {
		var err3 = devicesRef.Child(runtimeId).Set(ctx, dataOut)
		if err3 != nil {
			fmt.Println("Error with setting value of child: ", err3)
		}
	}()

}
