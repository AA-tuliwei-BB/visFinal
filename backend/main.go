package main

import (
	"backend/database"
	"backend/server"
	"fmt"
)

func main() {
	database.Open()
	defer database.Close()
	// database.PrintRecord(5)

	server.Start()

	fmt.Println("exit")
}
