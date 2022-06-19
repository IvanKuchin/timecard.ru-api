package apihandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ivankuchin/timecard.ru-api/logs"
)

func login_getHTTPPayload(u users) string {
	return "login=" + u.Login + "&password=" + u.Password
}

func login_convertRequest(tID string, body []byte) (string, error) {
	var user users

	err := json.Unmarshal(body, &user)
	if err != nil {
		error_message := "incorrect json format"
		logs.Sugar.Errorw(error_message+" (unmarshal error: "+err.Error()+")",
			"traceID", tID,
		)
		return "", fmt.Errorf("%s", error_message)
	}

	return login_getHTTPPayload(user), nil
}

func login_parseServerResponse(tID string, sr []byte) (string, error) {

	var server_response login_response
	err := json.Unmarshal(sr, &server_response)
	if err != nil {
		error_message := "incorrect json format"
		logs.Sugar.Errorw(error_message+" (unmarshal error: "+err.Error()+")",
			"traceID", tID,
		)
		return "", fmt.Errorf("%s", error_message)
	}

	return server_response.Sessid, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tID := generateTraceID()

	body, err := getClientRequestBody(tID, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	auth_params, err := login_convertRequest(tID, body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	url := config.Serverproto + "://" + config.Serverhost + ":" + strconv.Itoa(config.Serverport) + "/cgi-bin/noauth.cgi?action=API_login&" + auth_params
	server_response, err := sendReqToServerNoAuth(tID, url)
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

	sessid, err := login_parseServerResponse(tID, server_response)
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

	fmt.Fprintf(w, "%v", sessid)
}
