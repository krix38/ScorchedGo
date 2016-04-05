package rest

import (
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/krix38/ScorchedGo/properties"
)

type connectionInfo struct {
	signedIn bool
}

var ConnectionStatus = createRestHandler(connectionStatus, []string{"GET"})

func connectionStatus() interface{}{
		/* TODO: check connection status */
		return connectionInfo{signedIn: false}
}

func createRestHandler(handler func() interface{}, acceptedMethods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range acceptedMethods {
			if r.Method == method {
				w.Header().Set("Content-Type", "application/json")
				object := handler()
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
