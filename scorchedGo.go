package main

import (
	"log"

	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/controller"
)

func main() {
	log.Print(properties.Messages.DebugStarting + properties.Configuration.Port)
	log.Print(properties.Messages.DebugLanguageInfo)
	controller.RunWebController()
}
