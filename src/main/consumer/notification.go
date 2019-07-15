package main

import (
	"main/controllers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	controllers.ReceiveMessage()
}
