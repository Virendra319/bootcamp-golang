package ExercisesDay3

import (
	"awesomeProject/Config"
	"awesomeProject/Models"
	"awesomeProject/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func Problem2() {
	Config.DB, err = gorm.Open("mysql", Config.DBURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
