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
// swagger:meta
package apihandlers

// Returns list of bt invoices
type invoice_btServerResponse struct {
	// Result status
	// in: body
	Result string
	// bt invoices
	// in: body
	Bt_invoices []invoice
}

// Returns list of invoces
// swagger:response invoice_btServerResponseWrapper
type invoice_btServerResponseWrapper struct {
	// Result status
	// in: body
	Body invoice_btServerResponse
}

// Returns invoice details
type invoice_bt_detailServerResponse struct {
	// Result status
	// in: body
	Result string
	// List of timecatds
	// in: body
	Bt []bt
}

// Returns invoce details
// swagger:response invoice_bt_detailServerResponseWrapper
type invoice_bt_detailServerResponseWrapper struct {
	// Result status
	// in: body
	Body invoice_bt_detailServerResponse
}
