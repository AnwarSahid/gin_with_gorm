package main

import (
	"gin_with_gorm/database"
	"gin_with_gorm/routers"
)

func main() {
	database.StartDB()
	routers.StartServer().Run(":8080")

}
