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

// Service invoice structure
//
// swagger:model invoice
type invoice struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Reference to a cost center this invoice will be sent over
	// in: body
	//
	// required: true
	Cost_center_id string
	// URL to download full package, prefixed by /invoices_cc/
	// in: body
	//
	// required: true
	File string
	// Reference to an agency initianig invoicing
	// in: body
	//
	// required: true
	Owner_company_id string
	// Reference to a user generated invoice
	// in: body
	//
	// required: true
	Owner_user_id string
	// User object
	// in: body
	//
	// required: true
	Users []user
}

// Returns list of service invoices
type invoice_serviceServerResponse struct {
	// Result status
	// in: body
	Result string
	// Service invoices
	// in: body
	Service_invoices []invoice
}

// Returns list of invoces
// swagger:response invoice_serviceServerResponseWrapper
type invoice_serviceServerResponseWrapper struct {
	// Result status
	// in: body
	Body invoice_serviceServerResponse
}

// Returns invoice details
type invoice_service_detailServerResponse struct {
	// Result status
	// in: body
	Result string
	// List of timecatds
	// in: body
	Timecards []timecard
}

// Returns invoce details
// swagger:response invoice_service_detailServerResponseWrapper
type invoice_service_detailServerResponseWrapper struct {
	// Result status
	// in: body
	Body invoice_service_detailServerResponse
}
