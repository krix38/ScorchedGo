package httpErrorPrinter

import (
	"log"
	"net/http"
)

func Print(w http.ResponseWriter, r *http.Request, message string, httpcode int, logMessage string) {
	http.Error(w, message, httpcode)
	log.Println(r.RemoteAddr + ":" + logMessage)
}
