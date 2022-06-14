package apihandlers

import (
	"math/rand"

	configreader "github.com/ivankuchin/timecard.ru-api/config-reader"
	"go.uber.org/zap"
)

var config configreader.Config
var logger *zap.SugaredLogger

func generateTransID() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const transIDLen = 10

	transID := make([]byte, transIDLen)

	for i := range transID {
		transID[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(transID)
}

func SetConfig(c configreader.Config) {
	config = c
}

func SetLogger(l *zap.SugaredLogger) {
	logger = l
}
