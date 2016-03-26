package properties

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Port            string
	Language        string
	StaticFilesPath string
}

type messages struct {
	DebugStarting     string
	DebugLanguageInfo string
}

func GetConfiguration() *configuration {
	file := new(os.File)
	file, err := os.Open("conf/conf.json")
	if err != nil {
		file, err = os.Open("src/github.com/krix38/ScorchedGo/conf/conf.json")
		if err != nil {
			log.Fatal("error opening configuration file")
		}
	}
	decoder := json.NewDecoder(file)
	conf := new(configuration)
	err = decoder.Decode(conf)
	if err != nil {
		log.Fatal("error decoding json configuration file")
	}
	return conf
}

func GetMessages(conf *configuration) *messages {
	file := new(os.File)
	filename := "messages_" + conf.Language + ".json"
	file, err := os.Open("conf/" + filename)
	if err != nil {
		file, err = os.Open("src/github.com/krix38/ScorchedGo/conf/" + filename)
		if err != nil {
			log.Fatal("error opening " + filename + " file")
		}
	}
	decoder := json.NewDecoder(file)
	msg := new(messages)
	err = decoder.Decode(msg)
	if err != nil {
		log.Fatal("error decoding json " + filename + " file")
	}
	return msg
}

var Configuration *configuration = GetConfiguration()
var Messages *messages = GetMessages(Configuration)
