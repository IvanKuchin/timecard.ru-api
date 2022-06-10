package apihandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type users struct {
	login    string
	password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("ERROR: %f\n", err)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "not enough parameters\n")
		return
	}

	var user users

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "incorrect json format\n")
		return
	}

	fmt.Fprintln(w, "ok")
}
