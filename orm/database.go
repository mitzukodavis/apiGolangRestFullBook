package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitzukodavis/apirestgolang/config"
)

var db *gorm.DB


func CreateConnection()  {
	url := config.GetUrlDatabase()
	if connention , err := gorm.Open("mysql",url); err != nil{
		panic(err)
	}else {
		db = connention
	}
}


func CloseConnection()  {
	db.Close()
}

func CreateTables()  {
	db.DropTableIfExists(&User{})//
	db.CreateTable(&User{})
}

