package appruntime

import (
	"context"
	"fmt"
	"log"
	"proj-cluster/firebaseapp"
	"proj-cluster/heart"
	"proj-cluster/utils"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/google/uuid"
)

func RunNewRuntime(firebaseApp *firebase.App, ctx context.Context, firebaseDatabaseClient *db.Client) {
	runtimeId := uuid.New().String()
	asiaManilaLocation, locationError := time.LoadLocation("Asia/Manila")
	if locationError != nil {
		log.Fatalln("Error loading location: ", locationError)
	}
	startTimeValue := time.Now().In(asiaManilaLocation)

	for {
		heart.Heartbeat(runtimeId, ctx, firebaseDatabaseClient)
		utils.LogEvent(runtimeId, startTimeValue)

		devices, deviceRetrievalError := firebaseapp.GetDevices(ctx, firebaseDatabaseClient)
		if deviceRetrievalError != nil {
			utils.LogError(runtimeId, deviceRetrievalError)
			fmt.Printf("ERROR WITH DEVICE RETRIEVAL: %v", deviceRetrievalError.Error())
		}
		firebaseapp.FormatAndPrintDevices(devices)

		time.Sleep(10 * time.Second)
	}
}
