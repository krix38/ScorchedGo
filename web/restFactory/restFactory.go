package restFactory

import (
	"encoding/json"
	"net/http"

	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/httpErrorPrinter"
)

type RestHandlerData struct {
	writer       http.ResponseWriter
	reader       *http.Request
	inputJsonObj interface{}
}

func CreateRestHandler(handler func(*RestHandlerData) (interface{}, error), acceptedMethods []string, inputJson interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range acceptedMethods {
			if r.Method == method {
				restData, err := prepareHandler(w, r, inputJson)
				if err != nil {
					httpErrorPrinter.Print(w, r, "error parsing input json", http.StatusBadRequest, properties.Messages.DebugJsonParseFail)
				}
				object, err := handler(restData)
				if object != nil {
					jsonObject, err := json.Marshal(object)
					if err != nil {
						httpErrorPrinter.Print(w, r, "error parsing output json", http.StatusBadRequest, properties.Messages.DebugJsonParseFail)
					}
					w.Write(jsonObject)
					return
				} else {
					if err != nil {
						httpErrorPrinter.Print(w, r, "internal server error", http.StatusInternalServerError, properties.Messages.DebugInternalServerError)
					}
					http.Error(w, "", http.StatusOK)
				}
			}
		}
		http.Error(w, "bad request method: "+r.Method, http.StatusBadRequest)
	}
}

func prepareHandler(w http.ResponseWriter, r *http.Request, inputJson interface{}) (*RestHandlerData, error) {
	restData := RestHandlerData{writer: w, reader: r}
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
