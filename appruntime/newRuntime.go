package appruntime

import (
	"proj-cluster/heart"
	"proj-cluster/utils"
	"time"

	"github.com/google/uuid"
)

func RunNewRuntime(data) {
	runtimeId := uuid.New().String()
	startTimeValue := time.Now()

	for {
		heart.Heartbeat(runtimeId, ctx, databaseClient)
		utils.LogEvent(runtimeId, "Logged", startTimeValue)
		time.Sleep(10 * time.Second)
	}
}
