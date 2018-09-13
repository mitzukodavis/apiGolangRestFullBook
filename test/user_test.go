package test

import (
	"testing"
	"fmt"
	"math/rand"
	"github.com/mitzukodavis/apirestgolang/models"

	"time"
)

var user *models.User

const(
	id = 1
	username = "eduardo_gpg"
	password = "password"
	passwordHash ="$2a$10$Cp8UVJUEfWZPsP3wXm2Q8.CR2ad4vYET5eKvOuh7tUPnigJqJRB5i "
	email = "eduardo@codigo.com"
	createdDate = "2017-08-17"
)

func TestNewUser(t *testing.T)  {
	_, err := models.NewUser(username, password, email)
	if err != nil{
		t.Error("No es posible crear el objeto", err)
	}
}

func TestPassword(t *testing.T)  {
	user,_ := models.NewUser(username, password, email)
	if user.Password == password || len(user.Password) != 60{
		t.Error("no es posible cifrar el password")
	}
	}

func TestValidEmail(t *testing.T)  {
	if err := models.ValidEmail(email); err !=nil{
		t.Error("validacion erronea en el email")
	}
}

func TestInvalidEmail(t *testing.T)  {
	if err := models.ValidEmail("dsfsdfdsfdsf"); err == nil{
		t.Error("validacion erronea en el email")
	}
}

func TestUsernameLenght(t *testing.T)  {
	newUsername := username
	for i:=0; i<10; i++{
		newUsername += newUsername
	}
	_, err := models.NewUser(newUsername, password, email)
	if err == nil {
		t.Error("No es posible generar usuario con una username muy grande ")
	}
}

func TestLogin(t *testing.T)  {
	if valid := models.Login(username, password); !valid{
		t.Error("No es posible realizar el login")
	}
}

func TestNoLogin(t *testing.T)  {
	if valid := models.Login(randomUsername(), password); valid{
		t.Error("es posible realizar el login con parametros erroneos")
	}
}

func TestSave(t *testing.T)  {
	user,_:= models.NewUser(randomUsername(), password, email)
	if err := user.Save(); err !=nil{
		t.Error("no es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T)  {
	_, err := models.CreateUser(randomUsername(), password, email)
	if err != nil {
		t.Error("No es posible ")
	}
}

func TestUniqueUsername(t *testing.T)  {
	_, err := models.CreateUser(username, password, email)
	if err == nil{
		t.Error("es posible inser registos con usernames duplicados")
	}
}

func TestDuplicateUsername(t *testing.T)  {
	_, err := models.CreateUser(username, password, email)
	message := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'username'" ,username)
	if err.Error() !=message{
		t.Error("es posible tener un username en ela base de datos")
	}
}

func TestGetUser(t *testing.T)  {
	user := models.GetUserByID(id)
	t.Log(user.GetCreatedDate())
	if !equalsUser(user)|| !equalsCreateData(user.GetCreatedDate()){
		t.Error("noes posible tener el usuario")
	}
}

func TestGetUsers(t *testing.T)  {
	users := models.GetUsers()
	if len(users)==0{
		t.Error("no es posible obtener los usuarios")
	}
}

//func TestDeleteUser(t *testing.T)  {
//	if err := user.Delete(); err != nil {
//		t.Error("No es posible eliminar al usuario")
//	}
//}

func equalsCreateData(date time.Time)bool{
	t, _:= time.Parse("2006-01-02",createdDate)
	return t == date
	}

func equalsUser(user *models.User) bool {
	return user.Username == username && user.Email == email
}

func randomUsername() string {
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}