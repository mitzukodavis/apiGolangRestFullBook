package handlers

import (
	"net/http"
	"github.com/mitzukodavis/apirestgolang/utils"
	"fmt"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

func Authentication(function customeHandler)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		//aqui va nuestra logica

		if !utils.IsAuthenticated(r){
			http.Redirect(w,r,  "/users/login", http.StatusSeeOther)
			return
		}
		function(w, r)
	})
}

func MiddlewareTwo(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Este es el segundo Wrap!")
		handler.ServeHTTP(w, r)
	});
}