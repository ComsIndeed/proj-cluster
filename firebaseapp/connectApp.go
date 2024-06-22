package firebaseapp

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func ConnectToFirebase(databaseURL string) (*firebase.App, *db.Client) {

	// Options
	config := &firebase.Config{DatabaseURL: databaseURL}
	ctx := context.Background()
	option := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	// New App
	app, appInitializationError := firebase.NewApp(ctx, config, option)
	if appInitializationError != nil {
		log.Fatalln("App initialization error: ", appInitializationError)
	}

	// New Client
	databaseClient, databaseClientInitializationError := app.Database(ctx)
	if databaseClientInitializationError != nil {
		log.Fatalln("Database client initialization error: ", databaseClientInitializationError)
	}

	return app, databaseClient
}
