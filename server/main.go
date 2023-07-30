package main

import "react-flow-api/database"

func main() {
	db := database.NewDB()
	database.Migrate(db)
}
