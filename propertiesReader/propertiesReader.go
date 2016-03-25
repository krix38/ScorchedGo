package propertiesReader

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Port     string
	Language string
}

type Messages struct {
	DebugStarting     string
	DebugLanguageInfo string
}

func GetConfiguration() (*Configuration, error) {
	file := new(os.File)
	file, err := os.Open("conf/conf.json")
	if err != nil {
		file, err = os.Open("src/github.com/krix38/ScorchedGo/conf/conf.json")
		if err != nil {
			log.Print("error opening configuration file")
			return nil, err
		}
	}
	decoder := json.NewDecoder(file)
	configuration := new(Configuration)
	err = decoder.Decode(configuration)
	if err != nil {
		log.Print("error decoding json configuration file")
		return nil, err
	}
	return configuration, nil
}

func GetMessages(configuration *Configuration) (*Messages, error) {
	file := new(os.File)
	filename := "messages_" + configuration.Language + ".json"
	file, err := os.Open("conf/" + filename)
	if err != nil {
		file, err = os.Open("src/github.com/krix38/ScorchedGo/conf/" + filename)
		if err != nil {
			log.Print("error opening " + filename + " file")
			return nil, err
		}
	}
	decoder := json.NewDecoder(file)
	messages := new(Messages)
	err = decoder.Decode(messages)
	if err != nil {
		log.Print("error decoding json " + filename + " file")
		return nil, err
	}
	return messages, nil
}
