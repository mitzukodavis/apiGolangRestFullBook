package main

import (
	"github.com/mitzukodavis/apirestgolang/orm"
	"fmt"
)

func main() {
	orm.CreateConnection()
	orm.CreateTables()

	user := orm.NewUser("Eduardo", "123", "a@a.com")
	user.Save()

	users := orm.GetUsers()
	fmt.Println(users)

	user=orm.GetUser(1)
	user.Username = "codigofacilito"
	user.Password = "cambio de password"
	user.Email = "cambiodecorreo@a.com"
	user.Save()

	fmt.Println(user)

	user.Delete()

	orm.CloseConnection()

}