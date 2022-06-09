package main

import (
	"fmt"
	"log"

	"github.com/ivankuchin/timecard.ru-api/config_reader"
)

func main() {
	config, err := config_reader.Read()
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(config)
}
