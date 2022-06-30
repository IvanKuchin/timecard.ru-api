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

// Returns list of BT (business trips)
type subc_by_sowServerResponse struct {
	// Result status
	// in: body
	Result string
	// Company info
	// in: body
	Companies []company
}

// Returns list of business trips
// swagger:response subc_by_sowServerResponseWrapper
type subc_by_sowServerResponseWrapper struct {
	// Result status
	// in: body
	Body subc_by_sowServerResponse
}
