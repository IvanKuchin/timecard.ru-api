// Package handlers for the RESTful Server
//
// Documentation for REST API
//
//  Schemes: http, https
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

	"github.com/gorilla/mux"
	"github.com/ivankuchin/timecard.ru-api/logs"
)

func timecard_parseServerResponse(tID string, sr []byte) (*[]byte, error) {

	var server_response timecardServerResponse
	err := json.Unmarshal(sr, &server_response)
	if err != nil {
		error_message := "incorrect json format"
		logs.Sugar.Errorw(error_message+" (unmarshal error: "+err.Error()+")",
			"traceID", tID,
		)
		return nil, fmt.Errorf("%s", error_message)
	}

	if len(server_response.Timecards) == 0 {
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

// swagger:route GET /api/v1/agency/timecard/ Timecards noContentWrapper
// Return array of Timcards reported to agency
//
// Schemes: http, https
//
// Security:
//   api_key
//
// responses:
// 200: timecardServerResponseWrapper
// 404: notFoundWrapper
// 400: badRequestWrapper

// swagger:route GET /api/v1/agency/timecard/{id} Timecards idParamTimecard
// Return timecard with id
//
// Schemes: http, https
//
// Security:
//   api_key
//
// responses:
// 200: timecardServerResponseWrapper
// 404: notFoundWrapper
// 400: badRequestWrapper
func AgencyTimecardListHandler(w http.ResponseWriter, r *http.Request) {
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

	url := ""

	vars := mux.Vars(r)
	key, exists := vars["key"]
	if exists {
		url = config.Serverproto + "://" + config.Serverhost + ":" + strconv.Itoa(config.Serverport) + "/cgi-bin/agency.cgi?action=AJAX_getTimecardEntry&timecard_id=" + key
	} else {
		url = config.Serverproto + "://" + config.Serverhost + ":" + strconv.Itoa(config.Serverport) + "/cgi-bin/agency.cgi?action=AJAX_getTimecardList"
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

	responseToClient, err := timecard_parseServerResponse(tID, server_response)
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
