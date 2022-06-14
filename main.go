package main

import (
	"log"

	configreader "github.com/ivankuchin/timecard.ru-api/config-reader"
	"github.com/ivankuchin/timecard.ru-api/server"
)

func SetLogFlags() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	SetLogFlags()

	config, err := configreader.Read()
	if err != nil {
		log.Panic(err.Error())
	}

	server.SetConfig(*config)
	server.Run()

	log.Println("exit from program")
}
