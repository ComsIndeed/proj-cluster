package heart

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"firebase.google.com/go/db"

	utils "proj-cluster/utils"
)

func Heartbeat(runtimeId string, ctx context.Context, databaseClient *db.Client) (statusCode uint8) {
	var devicesRef = databaseClient.NewRef("deviceStatuses")

	var hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	var connectionInfo, err2 = utils.GetConnectionInfo()
	if err2 != nil {
		log.Fatalf("CONNECTION INFO ERROR: %v", err2)
	}

	fmt.Printf("CONNECTION INFO: %v\n", connectionInfo)

	var dataOut = map[string]interface{}{
		"id":              runtimeId,
		"deviceId":        hostname,
		"operatingSystem": runtime.GOOS,
		"connection": map[string]interface{}{
			"ssid":           connectionInfo.SSID,
			"carrier":        connectionInfo.Carrier,
			"signalStrength": connectionInfo.SignalStrength,
			"connectionType": connectionInfo.ConnectionType,
		},
		"time":     time.Now().Unix(),
		"services": []string{"PLACEHOLDER 1", "PLACEHOLDER 2"},
	}

	var err3 = devicesRef.Child(runtimeId).Set(ctx, dataOut)
	if err3 != nil {
		log.Fatalln("Error with setting value of child: ", err3)
		return 0
	}
	return 1
}
