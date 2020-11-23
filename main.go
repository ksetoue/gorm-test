//main.go
package main

import (
	"fmt"
	"gorm-test/Config"
	"gorm-test/Models"
	"gorm-test/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{}, &Models.Deal{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
