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

import "errors"

type users struct {
	Login    string
	Password string
}

type login_response struct {
	Result      string
	Description string
	Sessid      string
}

type sow_response struct {
	Result      string
	Description string
	Sow         string
}

var ErrorNoBearerToken = errors.New("no bearer token")
