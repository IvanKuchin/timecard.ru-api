package apihandlers

import (
	"fmt"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "thanks for visiting %v\nAPI documentation available here: /swagger\n", r.RequestURI)
}
