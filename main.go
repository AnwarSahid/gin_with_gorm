package main

import "gin_with_gorm/routers"

func main() {

	routers.StartServer().Run(":8080")

}
