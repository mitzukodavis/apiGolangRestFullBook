package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"./handlers/api/v1"
	"github.com/mitzukodavis/apirestgolang/config"
	"github.com/mitzukodavis/apirestgolang/handlers"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/users/new", handlers.NewUser).Methods("GET", "POST")
	mux.HandleFunc("/users/login", handlers.Login).Methods("GET")
	mux.HandleFunc("/users/logout", handlers.Logout).Methods("GET")

	editHandler := handlers.Authentication(handlers.UpdateUser)
	editHandler = handlers.MiddlewareTwo(editHandler)
	mux.Handle("/users/edit", editHandler).Methods("GET")



	mux.HandleFunc("/api/v1/users/", v1.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", v1.CreateUser).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.DeleteUser).Methods("DELETE")

	assets := http.FileServer(http.Dir("assets"))
	statics := http.StripPrefix("/assets", assets)
	mux.PathPrefix("/assets/").Handler(statics)

	log.Println("el servidor en el puerto ", config.ServerPort())
	server := &http.Server{
		Addr: config.UrlServer(),
		Handler : mux,
	}
	log.Fatal(server.ListenAndServe())
}

//http://localhost:8000/api/v1/users/2 -X PUT -d '{"username" : "cambiado", "pasword" : "cambiado" }' -H "Content-Type: aplication/json"
//http://localhost:8000/api/v1/users/1 -X DELETE
