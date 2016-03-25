package main

import (
	"log"

	"github.com/krix38/ScorchedGo/propertiesReader"
)

func main() {
	conf, err := propertiesReader.GetConfiguration()
	if err != nil {
		log.Panic("error reading configuartion, stopping application...")
	}
	messages, err := propertiesReader.GetMessages(conf)
	if err != nil {
		log.Panic("error reading messages, stopping application...")
	}
	log.Print(messages.DebugStarting + conf.Port)
	log.Print(messages.DebugLanguageInfo)
}
