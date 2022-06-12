package apihandlers

import configreader "github.com/ivankuchin/timecard.ru-api/config-reader"

var config configreader.Config

func SetConfig(c configreader.Config) {
	config = c
}
