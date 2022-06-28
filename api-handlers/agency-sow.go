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
// Security:
// - api_key:
//
// SecurityDefinitions:
// api_key:
//      type: apiKey
//      name: Authorization
//      in: header
//
// swagger:meta
package apihandlers

import (
	"encoding/json"
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

func sow_parseServerResponse(tID string, sr []byte) (*[]byte, error) {

	var server_response sowServerResponse
	err := json.Unmarshal(sr, &server_response)
	if err != nil {
		error_message := "incorrect json format"
		logs.Sugar.Errorw(error_message+" (unmarshal error: "+err.Error()+")",
			"traceID", tID,
		)
		return nil, fmt.Errorf("%s", error_message)
	}

	if len(server_response.Sow) == 0 {
		logs.Sugar.Debugw(ErrorNotFound.Error(),
			"traceID", tID,
		)
		return nil, ErrorNotFound
	}

	serialized, err := json.Marshal(server_response)
	if err != nil {
		error_message := "json marshaling error"
		logs.Sugar.Errorw(error_message+" (marshal error: "+err.Error()+")",
			"traceID", tID,
		)
		return nil, fmt.Errorf("%s", error_message)
	}

	return &serialized, nil
}

// swagger:route GET /api/v1/agency/sow/ Sow noContentWrapper
// Return array of StatementOfWorks with subcontractors
//
// Schemes: http, https
//
// Security:
//   api_key
//
// responses:
// 200: sowServerResponseWrapper
// 404: notFoundWrapper
// 400: badRequestWrapper

// swagger:route GET /api/v1/agency/sow/{id} Sow idParam
// Return subcontractor StatementOfWork with id
//
// Schemes: http, https
//
// Security:
//   api_key
//
// responses:
// 200: sowServerResponseWrapper
// 404: notFoundWrapper
// 400: badRequestWrapper
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

	responseToClient, err := sow_parseServerResponse(tID, server_response)
	if err != nil {
		if (err.Error() == "user not found") || (err.Error() == "You are not authorized") {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "%v", err)
		} else if errors.Is(err, ErrorNotFound) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%v", err)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%v", err)
		}
		return
	}

	fmt.Fprintf(w, "%s", *responseToClient)
}
