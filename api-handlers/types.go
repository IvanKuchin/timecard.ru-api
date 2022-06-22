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

type login_user struct {
	// User login
	// in: body
	// required: true
	// type: string
	Login string
	// User password
	// in: body
	// required: true
	// type: string
	Password string
}

// swagger:parameters loginID
type login_user_wrapper struct {
	// in: body
	Body login_user
}

type login_response struct {
	Result      string
	Description string
	Sessid      string
}

// Company structure
//
// swagger:model company
type company struct {
	// company ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
}

// Business trip template structure
//
// swagger:model bt_expense_template
type bt_expense_template struct {
	// template ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
}

// Company occupation
//
// swagger:model company_positions
type company_positions struct {
	// template ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
}

// User structure
//
// swagger:model user
type user struct {
	// user ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
}

// PSoW structure
//
// swagger:model psow
type psow struct {
	// SoW ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
}

// Task structure
//
// swagger:model task
type task struct {
	// Task ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
}

// SoW structure
//
// swagger:model sow
type sow struct {
	// SoW ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Agency holding the contract
	// in: body
	//
	// required: true
	Agency_company_id []company
	// URL to download agreement
	// in: body
	//
	// required: true
	Agreement_filename string
	// Array of persons to approve busines trips
	// in: body
	//
	// required: true
	Bt_approvers []user
	// Array of templates used for business trpi
	// in: body
	//
	// required: true
	Bt_expense_templates []bt_expense_template
	// Subcontractor occupation
	// in: body
	//
	// required: true
	Company_positions []company_positions
	// Cost center that will be charged for subcontractor service
	// in: body
	//
	// required: true
	Cost_centers []company
	// Dayrate with subcontractor
	// in: body
	//
	// required: true
	Day_rate string
	// Contract expiraton date
	// in: body
	//
	// required: true
	End_date string
	// Contract number (as written on paper)
	// in: body
	//
	// required: true
	Number string
	// Number of days to pay after business trip expense appruval
	// in: body
	//
	// required: true
	Payment_period_bt string
	// Number of days to pay after timecard appruval
	// in: body
	//
	// required: true
	Payment_period_service string
	// PSoW contract (Partner Statement of Work)
	// in: body
	//
	// required: true
	Psow []psow
	// Contract signing date
	// in: body
	//
	// required: true
	Sign_date string
	// Contract start date
	// in: body
	//
	// required: true
	Start_date string
	// Current contract status
	// in: body
	//
	// required: true
	Status string
	// Subcontractor company object
	// in: body
	//
	// required: true
	Subcontractor_company []company
	// Array of tasks assigned to contract
	// in: body
	//
	// required: true
	Tasks []task
	// Array of persons to approve timecards
	// in: body
	//
	// required: true
	Timecard_approvers []user
	// Reporting period (month or week)
	// in: body
	//
	// required: true
	Timecard_period string
	// Baseline of working day. Overtime calculaed based on this number.
	// in: body
	//
	// required: true
	Working_hours_per_day string
}

// Returns list of SoW (Statement of Works)
// swagger:response sowList
type sowContent struct {
	// Result status
	// in: body
	Result string
	// Array of SoW
	// in: body
	Sow []sow
}

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

// swagger:parameters idParam
type idParam struct {
	// The action id to be used in query
	// in: path
	// allowEmptyValue: true
	Id int `json:"id"`
}

var ErrorNoBearerToken = errors.New("no bearer token")
var ErrorNotFound = errors.New("not found")
