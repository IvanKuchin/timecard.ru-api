// Package handlers for the RESTful Server
//
// Documentation for REST API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.7
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta
package apihandlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ivankuchin/timecard.ru-api/logs"
)

func getBearerToken(tID string, r *http.Request) (string, error) {
	header_auth := r.Header.Get("Authorization")
	splitToken := strings.Split(header_auth, "Bearer ")

	if len(splitToken) == 1 {
		logs.Sugar.Errorw(ErrorNoBearerToken.Error(),
			"traceID", tID,
		)
		return "", ErrorNoBearerToken
	}
	token := splitToken[1]

	return token, nil
}

func AgencySowListHandler(w http.ResponseWriter, r *http.Request) {
	tID := generateTraceID()

	sessid, err := getBearerToken(tID, r)
	if err != nil {
		if errors.Is(err, ErrorNoBearerToken) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "%v", err)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%v", err)
		}
		return
	}

	url := config.Serverproto + "://" + config.Serverhost + ":" + strconv.Itoa(config.Serverport) + "/cgi-bin/agency.cgi?action=AJAX_getSoWList&include_bt=true&include_tasks=true"

	vars := mux.Vars(r)
	key, exists := vars["key"]
	if exists {
		url += "&sow_id=" + key
	}

	server_response, err := sendReqToServer(tID, url, sessid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	err = parseServerResponse(tID, server_response)
	if err != nil {
		if (err.Error() == "user not found") || (err.Error() == "You are not authorized") {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "%v", err)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%v", err)
		}
		return
	}

	fmt.Fprintf(w, "%s", server_response)
}