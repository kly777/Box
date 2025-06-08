package main

import (
	"fmt"

	"box/backend/database"
)

func main() {
	fmt.Println("Starting server...")
	database.InitDB()

}
