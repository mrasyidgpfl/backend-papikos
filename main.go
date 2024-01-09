package main

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/db"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/server"
	"fmt"
)

func main() {
	err := db.Connect()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	server.Init()
}
