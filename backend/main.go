package main

import (
	"backend/database"
	"fmt"
)

func main() {
	database.Open()
	defer database.Close()
	database.PrintRecord(5)
	fmt.Println("exit")
}
