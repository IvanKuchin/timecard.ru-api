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

// Return empty body
// swagger:response noContentWrapper
type noContentWrapper struct {
}

// Item not found
// swagger:response notFoundWrapper
type notFoundWrapper struct {
}

// Bad request
// swagger:response badRequestWrapper
type badRequestWrapper struct {
}

// Bearer token that must be used in future API-requests
// swagger:response bearerToken
type bearerToken struct {
	// Bearer token (for example: Bearer 1234567890)
	// in: body
	token string
}

// swagger:parameters idParamSoW idParamTimecard idParamBT idParamSubcBySow idParamInvoiceService
type idParam struct {
	// The action id to be used in query
	// in: path
	// allowEmptyValue: true
	Id int `json:"id"` // if json key is missed, request send from swagger doesn't assign id from <input>-form to the curl-request
}

var ErrorNoBearerToken = errors.New("no bearer token")
var ErrorNotFound = errors.New("not found")
