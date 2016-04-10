package session

import (
	"net/http"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("cookiestore"))

func CreateSessionHandler(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
}