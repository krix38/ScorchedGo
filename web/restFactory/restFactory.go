package restFactory

import (
	"encoding/json"
	"net/http"

	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/httpErrorPrinter"
)

type RestHandlerData struct {
	Writer       http.ResponseWriter
	Reader       *http.Request
	InputJsonObj interface{}
}

func CreateRestHandler(handler func(*RestHandlerData) (interface{}, error), acceptedMethods []string, inputJson interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buildRestHandler(w, r, handler, acceptedMethods, inputJson)
	}
}

func buildRestHandler(w http.ResponseWriter, r *http.Request, handler func(*RestHandlerData) (interface{}, error), acceptedMethods []string, inputJson interface{}) {
	createdHandler := false
	for _, method := range acceptedMethods {
		if r.Method == method {
			processRestHandler(w, r, handler, inputJson)
			createdHandler = true
		}
	}
	if !createdHandler {
		http.Error(w, "bad request method: "+r.Method, http.StatusBadRequest)
	}
}

func processRestHandler(w http.ResponseWriter, r *http.Request, handler func(*RestHandlerData) (interface{}, error), inputJson interface{}) {
	restData, err := prepareHandler(w, r, inputJson)
	if err != nil {
		httpErrorPrinter.Print(w, r, "error parsing input json", http.StatusBadRequest, properties.Messages.DebugJsonParseFail)
		return
	}
	object, err := handler(restData)
	if object != nil {
		jsonObject, err := json.Marshal(object)
		if err != nil {
			httpErrorPrinter.Print(w, r, "error parsing output json", http.StatusBadRequest, properties.Messages.DebugJsonParseFail)
			return
		}
		w.Write(jsonObject)
	} else {
		if err != nil {
			httpErrorPrinter.Print(w, r, "internal server error", http.StatusInternalServerError, err.Error())
			return
		}
		http.Error(w, "", http.StatusOK)
	}
}

func prepareHandler(w http.ResponseWriter, r *http.Request, inputJson interface{}) (*RestHandlerData, error) {
	restData := RestHandlerData{Writer: w, Reader: r}
	if inputJson != nil {
		err := decodeInputJson(inputJson, r)
		restData.InputJsonObj = inputJson
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
