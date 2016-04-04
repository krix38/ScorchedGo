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

func ConnectionStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		/* TODO: check connection status */
		connectionInfoStruct := connectionInfo{signedIn: false}
		connInfo, err := json.Marshal(connectionInfoStruct)
		if err != nil {
			log.Fatal(properties.Messages.DebugJsonParseFail + ": connectionInfoStruct")
		}
		w.Write(connInfo)
	} else {
		http.Error(w, properties.Messages.BadMethod + ": " + r.Method, http.StatusBadRequest)
	}
}
