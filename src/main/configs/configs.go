package configs

import (
	"fmt"
	"main/structs"

	"github.com/jinzhu/gorm"
)

func ConnectingToMySQL() {
	db, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("Connection Failed to open")
	} else {
		fmt.Println("Connection Established")
		if db.HasTable(&structs.NewsModel{}) {
			fmt.Println("Table News is exist")
		} else {
			db.AutoMigrate(&structs.NewsModel{})
			fmt.Println("Table News is created")
		}

	}
}
