package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/krix38/ScorchedGo/properties"
)

type restHandlerData struct {
	writer       http.ResponseWriter
	reader       *http.Request
	inputJsonObj interface{}
}

type connectionInfo struct {
	SignedIn bool
}

//
type testRest struct {
	TestVal string
}

var TestFunc = createRestHandler(testfunc, []string{"POST"}, testRest{})

func testfunc(restData *restHandlerData) (interface{}, error){
	return connectionInfo{SignedIn: true}, nil
}
//

var ConnectionStatus = createRestHandler(connectionStatus, []string{"GET"}, nil)

func connectionStatus(restData *restHandlerData) (interface{}, error) {
	/* TODO: check connection status */
	log.Print("inside connection status handler")
	return connectionInfo{SignedIn: false}, nil
}

func createRestHandler(handler func(*restHandlerData) (interface{}, error), acceptedMethods []string, inputJson interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range acceptedMethods {
			if r.Method == method {
				restData, err := prepareHandler(w, r, inputJson)
				if err != nil {
					httpErrorPrint(w, "error parsing input json", http.StatusBadRequest, properties.Messages.DebugJsonParseFail)
				}
				object, err := handler(restData)
				if object != nil {
					jsonObject, err := json.Marshal(object)
					if err != nil {
						httpErrorPrint(w, "error parsing output json", http.StatusBadRequest, properties.Messages.DebugJsonParseFail)
					}
					w.Write(jsonObject)
					return
				} else {
					if err != nil {
						httpErrorPrint(w, "internal server error", http.StatusInternalServerError, properties.Messages.DebugInternalServerError)
					}
					http.Error(w, "", http.StatusOK)
				}
			}
		}
		http.Error(w, "bad request method: "+r.Method, http.StatusBadRequest)
	}
}

func prepareHandler(w http.ResponseWriter, r *http.Request, inputJson interface{}) (*restHandlerData, error) {
	restData := restHandlerData{writer: w, reader: r}
	if inputJson != nil {
		err := decodeInputJson(inputJson, r)
		restData.inputJsonObj = inputJson
		if err != nil {
			return nil, err
		}
	}
	w.Header().Set("Content-Type", "application/json")
	return &restData, nil
}

func decodeInputJson(inputJson interface{}, r *http.Request) error {
	dec := json.NewDecoder(r.Body)
	return dec.Decode(&inputJson)
}

func httpErrorPrint(w http.ResponseWriter, message string, httpcode int, logMessage string) {
	http.Error(w, message, httpcode)
	log.Println(logMessage)
}
