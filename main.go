package main

import (
	"log"

	"github.com/ivankuchin/timecard.ru-api/config_reader"
	"github.com/ivankuchin/timecard.ru-api/server"
)

func main() {
	config, err := config_reader.Read()
	if err != nil {
		log.Panic(err.Error())
	}

	server.SetConfig(*config)
	server.Run()

	log.Println("exit from program")
}
