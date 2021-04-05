/*
	This is the starter function of the project.
	Connecting mysql.
	Auto migrating  modal and setup router form .
*/
package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)

		fmt.Println("connection failed to open")
	}

	fmt.Println("DB Conn Established")

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
