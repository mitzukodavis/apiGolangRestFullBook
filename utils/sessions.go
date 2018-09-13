package utils

import (
	"net/http"
	"github.com/satori/go.uuid"
	"time"
	"github.com/mitzukodavis/apirestgolang/models"
	"sync"
)

const(
	cookieName = "go_session"
	cookieExire = 24 * 2 * time.Hour //dos dias
)

var Sessions = struct {
	m map[string]*models.User
	sync.RWMutex
}{m: make(map[string]*models.User)}

func SetSession(user *models.User, w http.ResponseWriter)  {
	Sessions.Lock()
	defer Sessions.Unlock()
	uuid :=	uuid.NewV4().String()
	Sessions.m[uuid] = user
		cookie := &http.Cookie{
		Name: cookieName,
		Value: uuid,
		Path: "/",
		Expires:time.Now().Add(cookieExire),

	}
	http.SetCookie(w, cookie)
}

func GetUser(r *http.Request) *models.User  {
	Sessions.Lock()
	defer Sessions.Unlock()
	uuid := getValCookie(r)
	if user, ok := Sessions.m[uuid]; ok{
		return user
	}
	return &models.User{}
}

func DeleteSession(w http.ResponseWriter, r *http.Request)  {
	Sessions.Lock()
	defer Sessions.Unlock()

	delete(Sessions.m, getValCookie(r))
	cookie := &http.Cookie{
		Name: cookieName,
		Value: "",
		Path:"/",
		MaxAge:-1,
	}
	http.SetCookie(w, cookie)
}

func getValCookie(r *http.Request)  string{
	if cookie, err :=r.Cookie(cookieName); err == nil {
		return cookie.Value
	}
	return ""
}

func IsAuthenticated(r *http.Request) bool {
	return getValCookie(r) !=""
}