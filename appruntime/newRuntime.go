package appruntime

import (
	"context"
	"proj-cluster/heart"
	"proj-cluster/utils"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/google/uuid"
)

func RunNewRuntime(firebaseApp *firebase.App, ctx context.Context, firebaseDatabaseClient *db.Client) {
	runtimeId := uuid.New().String()
	startTimeValue := time.Now()

	for {
		heart.Heartbeat(runtimeId, ctx, firebaseDatabaseClient)
		utils.LogEvent(runtimeId, startTimeValue)
		time.Sleep(10 * time.Second)
	}
}
