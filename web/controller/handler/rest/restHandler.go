package rest

import (
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/krix38/ScorchedGo/properties"
)

type connectionInfo struct {
	SignedIn bool
}

var ConnectionStatus = createRestHandler(connectionStatus, []string{"GET"})

func connectionStatus(w http.ResponseWriter, r *http.Request) interface{}{
		/* TODO: check connection status */
		return connectionInfo{SignedIn: false}
}

func createRestHandler(handler func(w http.ResponseWriter, r *http.Request) interface{}, acceptedMethods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range acceptedMethods {
			if r.Method == method {
				w.Header().Set("Content-Type", "application/json")
				object := handler(w, r)
				jsonObject, err := json.Marshal(object)
				if err != nil {
					log.Fatal(properties.Messages.DebugJsonParseFail)
				}
				w.Write(jsonObject)
				return
			}
		}
		http.Error(w, properties.Messages.BadMethod+": "+r.Method, http.StatusBadRequest)
	}
}
