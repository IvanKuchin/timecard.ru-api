package apihandlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func readFromClient(tID string, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		logger.Errorw(err.Error(),
			"transactionID", tID,
		)
		return []byte{}, err
	}

	if len(body) == 0 {
		error_message := "not enough parameters\n"
		logger.Errorw(error_message,
			"transactionID", tID,
		)
		return []byte{}, fmt.Errorf("%s", error_message)
	}

	logger.Debugw("request url: "+r.RequestURI,
		"transactionID", tID,
	)
	logger.Debugw("request body: "+string(body),
		"transactionID", tID,
	)

	return body, nil
}

func getHTTPPayload(u users) string {
	return "login=" + u.Login + "&password=" + u.Password
}

func convertRequest(tID string, body []byte) (string, error) {
	var user users

	err := json.Unmarshal(body, &user)
	if err != nil {
		error_message := "incorrect json format"
		log.Printf("(tID: %s): ERROR: %s (unmarshal error: %v) (tried to parse: %s)\n", tID, error_message, err, body)
		return "", fmt.Errorf("%s", error_message)
	}

	return getHTTPPayload(user), nil
}

func sendReqToServer(tID string, url string) ([]byte, error) {
	buf := new(bytes.Buffer)

	req, err := http.NewRequest(http.MethodGet, url, buf)
	if err != nil {
		log.Printf("(tID: %s): ERROR: %v\n", tID, err)
		return []byte{}, fmt.Errorf("incorrect http request")
	}
	req.AddCookie(&http.Cookie{Name: "lng", Value: "us"})

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	c := &http.Client{
		Timeout: 5 * time.Second,
		// Transport: tr,
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Printf("(tID: %s): ERROR: %v\n", tID, err)
		return []byte{}, fmt.Errorf("error returned by server")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		error_message := fmt.Sprintf("server reply http.code %d\n", resp.StatusCode)
		log.Printf("(tID: %s) %s\n", tID, error_message)
		return []byte{}, fmt.Errorf("%s", error_message)
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("(tID: %s): ERROR: %v\n", tID, err)
		return []byte{}, fmt.Errorf("can't read from server")
	}

	return text, nil
}

func parseServerResponse(tID string, sr []byte) (string, error) {

	var server_response login_response
	err := json.Unmarshal(sr, &server_response)
	if err != nil {
		error_message := "incorrect json format"
		log.Printf("(tID: %s): ERROR: %s (unmarshal error: %v) (tried to parse: %s)\n", tID, error_message, err, string(sr))
		return "", fmt.Errorf("%s", error_message)
	}

	if server_response.Result == "error" {
		log.Printf("(tID: %s): ERROR: %s\n", tID, server_response.Description)
		return "", fmt.Errorf("%s", server_response.Description)
	}

	return server_response.Sessid, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tID := generateTransID()

	body, err := readFromClient(tID, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	auth_params, err := convertRequest(tID, body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	url := config.Serverproto + "://" + config.Serverhost + ":" + strconv.Itoa(config.Serverport) + "/cgi-bin/noauth.cgi?action=API_login&" + auth_params
	server_response, err := sendReqToServer(tID, url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	sessid, err := parseServerResponse(tID, server_response)
	if err != nil {
		if err.Error() == "user not found" {
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
