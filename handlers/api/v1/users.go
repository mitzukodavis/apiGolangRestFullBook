package v1

import (
	"net/http"
	"../../../models"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"errors"
)

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	//se genera un mapa
	//user := models.User{Id:1, Username: "eduardo_gpg", Password: "password123"}
	// se usa en el response //w.Header().Set("Content-Type", "application/json")
	//se necesita serializarlo para transportarlo por la red

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(w)
	}else {
		models.SendData(w, user)
	}
	//se elimina //response.Send()
	//se envio en response //output, _ := json.Marshal(&response)
	//fmt.Fprintf(w, string(output))
}

func CreateUser(w http.ResponseWriter, r *http.Request)  {
	user := &models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil{
		models.SendUnprocessableEntity(w)
		return
	}
	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	user.SetPassword(user.Password)
	if err := user.Save(); err !=nil{
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	user, err := getUserByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	request := &models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(request); err != nil{
		models.SendUnprocessableEntity(w)
		return
	}
	if err := user.Valid(); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	user.Username = request.Username
	user.Email = request.Email
	user.SetPassword(request.Password)

	if err := user.Save(); err !=nil{
		models.SendUnprocessableEntity(w)
		return
	}
	models.SendData(w, user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	user, err := getUserByRequest(r);
	if err != nil {
		models.SendNotFound(w)
	}else {
		user.Delete()
		models.SendNoContent(w)
	}
}

func getUserByRequest(r *http.Request) (*models.User, error)  {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi ( vars["id"] )
	user:= models.GetUserByID(userId)
	if user.Id == 0 {
		return user, errors.New("el usuario no existe en la DB")
	}
	return user, nil
}