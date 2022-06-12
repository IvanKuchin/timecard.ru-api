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

type users struct {
	Login    string
	Password string
}

func readFromClient(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("ERROR: %f\n", err)
		return []byte{}, err
	}

	if len(body) == 0 {
		error_message := "not enough parameters\n"
		log.Printf("ERROR: " + error_message)
		return []byte{}, fmt.Errorf("%s", error_message)
	}

	return body, nil
}

func getHTTPPayload(u users) string {
	return "login=" + u.Login + "&password=" + u.Password
}

func convertRequest(body []byte) (string, error) {
	var user users

	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return "", fmt.Errorf("incorrect json format")
	}

	return getHTTPPayload(user), nil
}

func sendReqToServer(url string) ([]byte, error) {
	location := config.Serverproto + "://" + config.Serverhost + ":" + strconv.Itoa(config.Serverport) + "/cgi-bin/index.cgi?action=AJAX_loginUser&" + url
	buf := new(bytes.Buffer)

	log.Printf("%q\n", location)

	req, err := http.NewRequest(http.MethodGet, location, buf)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return []byte{}, fmt.Errorf("incorrect http request")
	}

	// tr := &http.Transport{
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	c := &http.Client{
		Timeout: 5 * time.Second,
		// Transport: tr,
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return []byte{}, fmt.Errorf("error returned by server")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		error_message := fmt.Sprintf("server reply http.code %d\n", resp.StatusCode)
		log.Printf("%s\n", error_message)
		return []byte{}, fmt.Errorf("%s", error_message)
	}

	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return []byte{}, fmt.Errorf("can't read from server")
	}

	return text, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := readFromClient(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	request_url, err := convertRequest(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	server_response, err := sendReqToServer(request_url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	fmt.Fprintf(w, "ok\n%v\n", string(server_response))
}
