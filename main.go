package main

import (
	appruntime "proj-cluster/appruntime"
	firebaseapp "proj-cluster/firebaseapp"
)

func main() {
	app, databaseClient := firebaseapp.ConnectToFirebase(
		"https://proyekto-kumpol-default-rtdb.firebaseio.com/",
	)

	appruntime.RunNewRuntime({app: app, datadatabaseClient: datadatabaseClient})
	
}