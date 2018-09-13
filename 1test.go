package main

import (
	"github.com/mitzukodavis/apirestgolang/models"
	"fmt"
)

func main() {
	models.CreateConnection()
	models.CreateTables()



	models.CreateUser("Eduardo1", "123", "a@a.com")
	models.CreateUser("Eduardo2", "123", "a@a.com")
	models.CreateUser("Eduardo3", "123", "a@a.com")
	//
	//
	//fmt.Println(user)
	//
	//user.Username = "cambio de nombre"
	//user.Password = "cambio de password"
	//user.Email = "cambio de email"
	//user.Save()
	//
	//user.Delete()
	users := models.GetUsers()
	fmt.Println(users)

	models.CloseConnection()

}