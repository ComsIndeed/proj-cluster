package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"proj-cluster/heart"
	"time"

	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

func main() {
	var runtimeId = uuid.New().String()
	fmt.Println(runtimeId)

	var configuration = &firebase.Config{
		DatabaseURL: "https://proyekto-kumpol-default-rtdb.firebaseio.com/",
	}

	var option = option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	var ctx = context.Background()
	var app, appInitializationError = firebase.NewApp(ctx, configuration, option)
	if appInitializationError != nil {
		log.Fatalln("App initialization error: ", appInitializationError)
	}

	var databaseClient, databaseClientInitializationError = app.Database(ctx)
	if databaseClientInitializationError != nil {
		log.Fatalln("Database client initialization error: ", databaseClientInitializationError)
	}

	fmt.Printf("Running as %s", runtimeId)
	for {
		heart.Heartbeat(runtimeId, ctx, databaseClient)
		time.Sleep(10 * time.Second)
	}
}
