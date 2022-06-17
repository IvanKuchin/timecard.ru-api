package apihandlers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	configreader "github.com/ivankuchin/timecard.ru-api/config-reader"
	"github.com/ivankuchin/timecard.ru-api/logs"
)

var config configreader.Config

func SetConfig(c configreader.Config) {
	config = c
}

func generateTransID() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const transIDLen = 10

	transID := make([]byte, transIDLen)

	for i := range transID {
		transID[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(transID)
}

func getClientRequestBody(tID string, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		logs.Sugar.Errorw(err.Error(),
			"transactionID", tID,
		)
		return []byte{}, err
	}

	if len(body) == 0 {
		error_message := "not enough parameters\n"
		logs.Sugar.Errorw(error_message,
			"transactionID", tID,
		)
		return []byte{}, fmt.Errorf("%s", error_message)
	}

	logs.Sugar.Debugw("request url: "+r.RequestURI,
		"transactionID", tID,
	)
	logs.Sugar.Debugw("request body: "+string(body),
		"transactionID", tID,
	)

	return body, nil
}
