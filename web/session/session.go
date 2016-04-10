package session

import (
	"net/http"
)

func CreateSessionHandler(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
}
