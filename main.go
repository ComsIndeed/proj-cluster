package main

import (
	"context"
	appruntime "proj-cluster/appruntime"
	firebaseapp "proj-cluster/firebaseapp"
)

func main() {

	// Login and connect to Firebase
	ctx := context.Background()
	app, databaseClient := firebaseapp.ConnectToFirebase(
		ctx,
		"https://termux-server-cluster-default-rtdb.firebaseio.com/",
	)

	// Run a new instance of the thing sending heartbeats
	appruntime.RunNewRuntime(app, ctx, databaseClient)

}

// TODO
// TODO		Make it so that when it loses connection to Firebase, it reattempts connection
// TODO
// TODO		Make it so that the logs state the status of both the runtime, firebase connection, and the services
// TODO
